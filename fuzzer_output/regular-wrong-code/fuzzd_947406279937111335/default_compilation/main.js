// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(855),new _dafny.CodePoint('c'.codePointAt(0)))).length), new BigNumber(-525), (_dafny.ZERO).minus(new BigNumber(-448)))).Union(_dafny.Set.fromElements(new BigNumber(289), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, false, !(true), false, true)).length))).length)))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dfanxb")).length), new BigNumber(-110), new BigNumber(-828)));
    };
    static fm1(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('i'.codePointAt(0));
    };
    static fm2(p0, globalState) {
      return (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true)).Keys.Elements) {
          let _0_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true)).contains(_0_v0)) {
            _coll0.push([_0_v0,false]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),true));
    };
    static fm3(globalState) {
      return (_dafny.ZERO).minus((_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false, false)).length))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(406), new BigNumber(244), new BigNumber(694), new BigNumber((_dafny.Seq.UnicodeFromString("tudpflvdl")).length))).length))).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(329),new BigNumber((_dafny.Seq.UnicodeFromString("nmmeb")).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("rdcukv")).length))));
    };
    static fm4(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!((_dafny.Set.fromElements(new BigNumber(831), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(251))).length))).cardinality()))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("famu")).length), new BigNumber(((_module.D3.create_DC8(_dafny.Seq.UnicodeFromString("fsdvonkfs"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_1_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("w")).length);
}))).dtor_cf13).length)))),true);
    };
    static fm5(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(!(false))).Union(_dafny.MultiSet.fromElements(!(true), false, !(true)))).Union((_dafny.MultiSet.fromElements(true, true, false, true, false)).Intersect(_dafny.MultiSet.fromElements(false)));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC8(_dafny.Seq.UnicodeFromString("jucvf"), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), function (_2_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(604), new BigNumber(866)),new _dafny.CodePoint('u'.codePointAt(0)))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(512)), function (_3_i1) {
  return new BigNumber((_dafny.Seq.of(!(true))).length);
})));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), function (_4_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }),_dafny.Seq.UnicodeFromString("bxf"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("epp"),_dafny.Seq.UnicodeFromString("kjb")))).Merge((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("b"),!(false))).Keys.Elements) {
          let _5_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("b"),!(false))).contains(_5_v0)) {
            _coll1.push([_5_v0,_dafny.Seq.UnicodeFromString("lcmexjuub")]);
          }
        }
        return _coll1;
      }()).Merge((_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-223)), function (_6_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}),_dafny.Seq.UnicodeFromString("lsqfg")))).dtor_cf29));
    };
    static fm8(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("obosvty"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_7_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })),false);
    };
    static fm9(p0, p1, globalState) {
      return _module.D3.create_DC7((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true)));
    };
    static fm10(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-538)), function (_8_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), function (_9_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("fn")));
    };
    static fm11(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(false))).length))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length))),((false) ? (new BigNumber(-84)) : (new BigNumber(-734))));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _hi0 = (p3).plus(p0);
      for (let _10_i0 = p3; _10_i0.isLessThan(_hi0); _10_i0 = _10_i0.plus(_dafny.ONE)) {
        let _source0 = _module.D2.create_DC6(!(_10_i0).isEqualTo(p0));
        if (_source0.is_DC6) {
          let _11___mcc_h0 = (_source0).cf10;
          let _12_cf10 = _11___mcc_h0;
          let _13_v0;
          let _nw0 = new _module.C0();
          _nw0.__ctor(p0);
          _13_v0 = _nw0;
          let _14_v1;
          let _nw1 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _14_v1 = _nw1;
          let _index0 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_14_v1).length));
          (_14_v1)[_index0] = new BigNumber((_module.__default.fm10(!(p2), globalState)).length);
          let _index1 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_14_v1).length));
          (_14_v1)[_index1] = _module.__default.safeModuloInt((p0).plus(_10_i0), p3);
          let _15_v2;
          _15_v2 = _dafny.Seq.of(p2);
          _15_v2 = _15_v2;
          let _index2 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_14_v1).length));
          (_14_v1)[_index2] = new BigNumber(-844);
        } else {
          let _16___mcc_h1 = (_source0).cf9;
          let _17_cf9 = _16___mcc_h1;
          let _18_v3;
          let _nw2 = Array((new BigNumber(9)).toNumber()).fill([]);
          _18_v3 = _nw2;
          let _19_v4;
          let _nw3 = Array((new BigNumber(18)).toNumber()).fill(false);
          _19_v4 = _nw3;
          let _index3 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_18_v3).length));
          (_18_v3)[_index3] = _19_v4;
          let _20_v5;
          let _init0 = function (_21_i1) {
            return _module.__default.safeModuloInt(_21_i1, new BigNumber(805));
          };
          let _nw4 = Array((new BigNumber(3)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
            _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _20_v5 = _nw4;
          let _index4 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_20_v5).length));
          (_20_v5)[_index4] = _10_i0;
          let _22_v6;
          _22_v6 = _dafny.Seq.UnicodeFromString("cfaaiqx");
          let _23_v7;
          _23_v7 = _dafny.Map.Empty.slice().updateUnsafe((((_17_cf9)[_module.__default.safeIndex(_10_i0, new BigNumber((_17_cf9).length))]) ? (p2) : (p2)),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ma"), _22_v6));
          let _24_v8;
          _24_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _25_v9;
          _25_v9 = new _dafny.CodePoint('p'.codePointAt(0));
          let _26_v10;
          _26_v10 = _dafny.Set.fromElements(p2);
          let _index5 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_18_v3).length));
          let _index6 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_20_v5).length));
          let _rhs0 = _dafny.Seq.update(_dafny.Seq.update((((_23_v7).contains((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) ? ((_23_v7).get((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) : (_22_v6)), _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p3)), new BigNumber(((((_23_v7).contains((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) ? ((_23_v7).get((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) : (_22_v6))).length)), _25_v9), _module.__default.safeIndex(_module.__default.safeModuloInt(_module.__default.fm3(globalState), new BigNumber((_22_v6).length)), new BigNumber((_dafny.Seq.update((((_23_v7).contains((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) ? ((_23_v7).get((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) : (_22_v6)), _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p3)), new BigNumber(((((_23_v7).contains((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) ? ((_23_v7).get((((_24_v8).contains(p2)) ? ((_24_v8).get(p2)) : (p2)))) : (_22_v6))).length)), _25_v9)).length)), _25_v9);
          let _rhs1 = _19_v4;
          let _rhs2 = _10_i0;
          let _rhs3 = ((p2) ? (((p2) ? (p3) : (new BigNumber((_26_v10).length)))) : (_10_i0));
          let _lhs0 = globalState;
          let _lhs1 = _18_v3;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_18_v3).length));
          let _lhs3 = _20_v5;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_20_v5).length));
          let _lhs5 = globalState;
          _lhs0.f4 = _rhs0;
          _lhs1[_lhs2] = _rhs1;
          _lhs3[_lhs4] = _rhs2;
          _lhs5.f7 = _rhs3;
          let _index7 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_19_v4).length));
          (_19_v4)[_index7] = p2;
          let _index8 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_19_v4).length));
          (_19_v4)[_index8] = ((false) ? (p2) : (p2));
          let _27_v11;
          _27_v11 = _dafny.Map.Empty.slice().updateUnsafe(_25_v9,p3);
          let _28_v12;
          _28_v12 = _module.D0.create_DC2(p0);
          let _29_v13;
          _29_v13 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _30_v14;
          _30_v14 = _dafny.Set.fromElements(_28_v12, _28_v12, _module.D0.create_DC2((_20_v5)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_20_v5).length))]), _module.D0.create_DC2(new BigNumber((_29_v13).length)), _28_v12);
          let _index9 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_19_v4).length));
          let _rhs4 = _module.__default.safeModuloInt(new BigNumber((_27_v11).length), (new BigNumber((_30_v14).length)).plus(p0));
          let _rhs5 = p2;
          let _rhs6 = _dafny.Seq.Concat(_22_v6, _22_v6);
          let _lhs6 = globalState;
          let _lhs7 = _19_v4;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_19_v4).length));
          let _lhs9 = globalState;
          _lhs6.f7 = _rhs4;
          _lhs7[_lhs8] = _rhs5;
          _lhs9.f4 = _rhs6;
          let _31_v15;
          _31_v15 = _dafny.Map.Empty.slice().updateUnsafe(!((_19_v4)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_19_v4).length))]),(_18_v3)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_18_v3).length))]);
          let _32_v16;
          _32_v16 = _dafny.Map.Empty.slice().updateUnsafe((_31_v15).Merge(_31_v15),_module.__default.safeModuloInt((_20_v5)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_20_v5).length))], p3));
          _32_v16 = (_32_v16).update((_31_v15).update(false, (_18_v3)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_18_v3).length))]), p3);
        }
        let _33_v17;
        _33_v17 = _module.D0.create_DC2(new BigNumber(993));
        let _source1 = function (_pat_let0_0) {
          return function (_34_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_36_dt__update_hcf4_h0) {
                return _module.D0.create_DC2(_36_dt__update_hcf4_h0);
              }(_pat_let1_0);
            }(new BigNumber((function () {
              let _coll2 = new _dafny.Map();
              for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-54), new BigNumber(138))) {
                let _35_v18 = _compr_2;
                if (((new BigNumber(-54)).isLessThanOrEqualTo(_35_v18)) && ((_35_v18).isLessThan(new BigNumber(138)))) {
                  _coll2.push([(_35_v18).plus(new BigNumber((_dafny.Seq.UnicodeFromString("fqe")).length)),new BigNumber(776)]);
                }
              }
              return _coll2;
            }()).length));
          }(_pat_let0_0);
        }(_33_v17);
        if (_source1.is_DC1) {
          let _37___mcc_h2 = (_source1).cf1;
          let _38___mcc_h3 = (_source1).cf2;
          let _39___mcc_h4 = (_source1).cf3;
          let _40_cf3 = _39___mcc_h4;
          let _41_cf2 = _38___mcc_h3;
          let _42_cf1 = _37___mcc_h2;
          let _43_v19;
          let _nw5 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _43_v19 = _nw5;
          let _44_v20;
          let _nw6 = new _module.C0();
          _nw6.__ctor(p3);
          _44_v20 = _nw6;
          let _45_v21;
          _45_v21 = _dafny.Seq.of(_44_v20, _44_v20);
          let _index10 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_43_v19).length));
          (_43_v19)[_index10] = _dafny.Seq.Concat(_45_v21, _45_v21);
          let _index11 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_43_v19).length));
          (_43_v19)[_index11] = _dafny.Seq.of(_44_v20);
          _40_cf3 = _module.__default.safeModuloInt(p0, _40_cf3);
          _40_cf3 = _module.__default.safeModuloInt(_42_cf1, _10_i0);
          let _46_v22;
          _46_v22 = _dafny.Seq.UnicodeFromString("s");
          let _47_v23;
          _47_v23 = new _dafny.CodePoint('u'.codePointAt(0));
          (globalState).f4 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_46_v22, _46_v22), _module.__default.safeIndex(_42_cf1, new BigNumber((_dafny.Seq.Concat(_46_v22, _46_v22)).length)), _47_v23), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ihfuqk"), _dafny.Seq.UnicodeFromString("unigk")));
        } else if (_source1.is_DC2) {
          let _48___mcc_h5 = (_source1).cf4;
          let _49_cf4 = _48___mcc_h5;
          let _50_v24;
          let _nw7 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
          _50_v24 = _nw7;
          let _51_v25;
          _51_v25 = _dafny.Seq.UnicodeFromString("vp");
          let _52_v26;
          _52_v26 = _dafny.Seq.of(_51_v25);
          let _index12 = _module.__default.safeIndex(new BigNumber(444), new BigNumber((_50_v24).length));
          (_50_v24)[_index12] = _52_v26;
          let _index13 = _module.__default.safeIndex(new BigNumber(444), new BigNumber((_50_v24).length));
          (_50_v24)[_index13] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_53_v25) => function (_54_i2) {
            return _53_v25;
          })(_51_v25)), _dafny.Seq.Concat(_52_v26, _52_v26));
          let _55_v27;
          let _init1 = ((_56_cf4) => function (_57_i3) {
            return _module.__default.safeModuloInt(_57_i3, _56_cf4);
          })(_49_cf4);
          let _nw8 = Array((new BigNumber(21)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
            _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _55_v27 = _nw8;
          let _index14 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_55_v27).length));
          (_55_v27)[_index14] = p0;
          let _index15 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_55_v27).length));
          (_55_v27)[_index15] = _10_i0;
          let _58_v28;
          _58_v28 = false;
          _58_v28 = _58_v28;
          _58_v28 = _58_v28;
        } else if (_source1.is_DC3) {
          let _59___mcc_h6 = (_source1).cf5;
          let _60___mcc_h7 = (_source1).cf6;
          let _61___mcc_h8 = (_source1).cf7;
          let _62_cf7 = _61___mcc_h8;
          let _63_cf6 = _60___mcc_h7;
          let _64_cf5 = _59___mcc_h6;
          (globalState).f7 = p3;
          _63_cf6 = _63_cf6;
          let _65_v29;
          let _nw9 = new _module.C0();
          _nw9.__ctor(new BigNumber(-122));
          _65_v29 = _nw9;
          (_65_v29).f8 = _62_cf7;
        } else {
          let _66___mcc_h9 = (_source1).cf0;
          let _67_cf0 = _66___mcc_h9;
          let _68_v30;
          let _nw10 = new _module.C0();
          _nw10.__ctor(p0);
          _68_v30 = _nw10;
          let _69_v31;
          let _nw11 = Array((_dafny.ONE).toNumber()).fill(false);
          _69_v31 = _nw11;
          let _70_v32;
          _70_v32 = _dafny.Seq.of(_69_v31, _69_v31, _69_v31);
          _70_v32 = _70_v32;
          let _71_v33;
          _71_v33 = _dafny.Map.Empty.slice().updateUnsafe(_68_v30.f8,!(p3).isEqualTo(_10_i0));
          _71_v33 = (_71_v33).update(p3, p2);
          let _72_v34;
          _72_v34 = true;
          _72_v34 = _72_v34;
        }
        let _73_v35;
        _73_v35 = false;
        let _74_v36;
        _74_v36 = _dafny.Seq.UnicodeFromString("c");
        _73_v35 = (new BigNumber((_dafny.Seq.Concat(_74_v36, _dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), function (_75_i4) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }))).length)).isLessThanOrEqualTo(p0);
        let _76_v37;
        let _init2 = function (_77_i5) {
          return (_77_i5).minus(new BigNumber(-109));
        };
        let _nw12 = Array((new BigNumber(9)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw12.length); _i0_2++) {
          _nw12[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _76_v37 = _nw12;
        let _index16 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_76_v37).length));
        (_76_v37)[_index16] = (p3).minus(p3);
        let _index17 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_76_v37).length));
        (_76_v37)[_index17] = p0;
      }
      let _78_i6;
      _78_i6 = _dafny.ZERO;
      L0: {
        while (p2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_78_i6)) {
              break L0;
            }
            _78_i6 = (_78_i6).plus(_dafny.ONE);
            let _79_v38;
            _79_v38 = new _dafny.CodePoint('k'.codePointAt(0));
            let _80_v39;
            _80_v39 = _dafny.Map.Empty.slice().updateUnsafe(_79_v38,p2);
            let _81_v40;
            _81_v40 = _dafny.MultiSet.fromElements(_80_v39, _80_v39);
            let _82_v41;
            _82_v41 = _module.D0.create_DC3(_module.__default.fm11(globalState), _module.__default.fm0(!(p2), p2, _81_v40, p0, globalState), p0);
            (globalState).f7 = (_82_v41).dtor_cf7;
            let _83_v42;
            let _nw13 = new _module.C0();
            _nw13.__ctor(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), ((_84_v38) => function (_85_i7) {
              return _84_v38;
            })(_79_v38))).length));
            _83_v42 = _nw13;
            let _86_v43;
            _86_v43 = _module.D4.create_DC10(p2, _83_v42);
            let _pat_let_tv0 = p2;
            let _pat_let_tv1 = _83_v42;
            let _87_v44;
            let _nw14 = Array((new BigNumber(29)).toNumber());
            _nw14[(_dafny.ZERO).toNumber()] = _86_v43;
            _nw14[(_dafny.ONE).toNumber()] = function (_pat_let2_0) {
              return function (_88_dt__update__tmp_h1) {
                return function (_pat_let3_0) {
                  return function (_89_dt__update_hcf15_h0) {
                    return _module.D4.create_DC10(_89_dt__update_hcf15_h0, (_88_dt__update__tmp_h1).dtor_cf16);
                  }(_pat_let3_0);
                }(_pat_let_tv0);
              }(_pat_let2_0);
            }(_86_v43);
            _nw14[(new BigNumber(2)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(3)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(4)).toNumber()] = _module.D4.create_DC10(p2, _83_v42);
            _nw14[(new BigNumber(5)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(6)).toNumber()] = function (_pat_let4_0) {
              return function (_90_dt__update__tmp_h2) {
                return function (_pat_let5_0) {
                  return function (_91_dt__update_hcf16_h0) {
                    return _module.D4.create_DC10((_90_dt__update__tmp_h2).dtor_cf15, _91_dt__update_hcf16_h0);
                  }(_pat_let5_0);
                }(_pat_let_tv1);
              }(_pat_let4_0);
            }(_86_v43);
            _nw14[(new BigNumber(7)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(8)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(9)).toNumber()] = _module.D4.create_DC10(p2, _83_v42);
            _nw14[(new BigNumber(10)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(11)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(12)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(13)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(14)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(15)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(16)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(17)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(18)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(19)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(20)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(21)).toNumber()] = _module.D4.create_DC10(p2, _83_v42);
            _nw14[(new BigNumber(22)).toNumber()] = _module.D4.create_DC10(p2, _83_v42);
            _nw14[(new BigNumber(23)).toNumber()] = _module.D4.create_DC10(!(p2), _83_v42);
            _nw14[(new BigNumber(24)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(25)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(26)).toNumber()] = _module.D4.create_DC10(p2, _83_v42);
            _nw14[(new BigNumber(27)).toNumber()] = _86_v43;
            _nw14[(new BigNumber(28)).toNumber()] = _86_v43;
            _87_v44 = _nw14;
            let _92_v45;
            let _nw15 = Array((new BigNumber(17)).toNumber());
            _nw15[(_dafny.ZERO).toNumber()] = _87_v44;
            _nw15[(_dafny.ONE).toNumber()] = _87_v44;
            _nw15[(new BigNumber(2)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(3)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(4)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(5)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(6)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(7)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(8)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(9)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(10)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(11)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(12)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(13)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(14)).toNumber()] = _87_v44;
            _nw15[(new BigNumber(15)).toNumber()] = ((p2) ? (_87_v44) : (_87_v44));
            _nw15[(new BigNumber(16)).toNumber()] = _87_v44;
            _92_v45 = _nw15;
            let _index18 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_92_v45).length));
            (_92_v45)[_index18] = _87_v44;
            let _pat_let_tv2 = p2;
            let _index19 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_92_v45).length));
            let _nw16 = Array((new BigNumber(2)).toNumber());
            _nw16[(_dafny.ZERO).toNumber()] = function (_pat_let6_0) {
              return function (_93_dt__update__tmp_h3) {
                return function (_pat_let7_0) {
                  return function (_94_dt__update_hcf15_h1) {
                    return _module.D4.create_DC10(_94_dt__update_hcf15_h1, (_93_dt__update__tmp_h3).dtor_cf16);
                  }(_pat_let7_0);
                }(_pat_let_tv2);
              }(_pat_let6_0);
            }(_86_v43);
            _nw16[(_dafny.ONE).toNumber()] = _86_v43;
            (_92_v45)[_index19] = _nw16;
            let _95_v46;
            let _nw17 = new _module.C0();
            _nw17.__ctor(p0);
            _95_v46 = _nw17;
            let _96_v47;
            _96_v47 = _dafny.Seq.UnicodeFromString("ogpsjgmip");
            let _97_v48;
            _97_v48 = _dafny.Map.Empty.slice().updateUnsafe(_83_v42,_96_v47);
            _97_v48 = (_97_v48).update(_95_v46, _96_v47);
          }
        }
      }
      let _98_v49;
      let _nw18 = new _module.C0();
      _nw18.__ctor(p3);
      _98_v49 = _nw18;
      let _99_v50;
      _99_v50 = new _dafny.CodePoint('s'.codePointAt(0));
      let _100_v51;
      _100_v51 = _dafny.Map.Empty.slice().updateUnsafe(_98_v49,_99_v50);
      let _101_v53;
      let _init3 = ((_102_p0) => function (_103_i8) {
        return (_103_i8).multipliedBy((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.Set.fromElements(_102_p0)).Elements) {
            let _104_v52 = _compr_3;
            if ((_dafny.Set.fromElements(_102_p0)).contains(_104_v52)) {
              _coll3.add((_104_v52).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)));
            }
          }
          return _coll3;
        }()).length)));
      })(p0);
      let _nw19 = Array((new BigNumber(13)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw19.length); _i0_3++) {
        _nw19[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _101_v53 = _nw19;
      let _index20 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_101_v53).length));
      (_101_v53)[_index20] = (_98_v49.f8).minus((_dafny.ZERO).minus(p3));
      let _105_v54;
      _105_v54 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),p2);
      let _index21 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_101_v53).length));
      let _rhs7 = _dafny.Map.Empty.slice().updateUnsafe(_98_v49,_99_v50);
      let _rhs8 = p0;
      let _rhs9 = (new BigNumber((_105_v54).length)).multipliedBy((p3).minus(_98_v49.f8));
      let _rhs10 = _98_v49;
      let _lhs10 = _101_v53;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_101_v53).length));
      let _lhs12 = globalState;
      _100_v51 = _rhs7;
      _lhs10[_lhs11] = _rhs8;
      _lhs12.f7 = _rhs9;
      _98_v49 = _rhs10;
      let _106_v55;
      let _nw20 = Array((new BigNumber(10)).toNumber()).fill(false);
      _106_v55 = _nw20;
      _106_v55 = _106_v55;
      let _107_v56;
      _107_v56 = _dafny.Seq.UnicodeFromString("skwswvs");
      let _108_v57;
      _108_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(973),_107_v56);
      let _109_v58;
      _109_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Concat(_107_v56, (((_108_v57).contains(_98_v49.f8)) ? ((_108_v57).get(_98_v49.f8)) : (_dafny.Seq.UnicodeFromString("kefocyf"))))).length));
      let _110_v59;
      _110_v59 = _dafny.Map.Empty.slice().updateUnsafe(_99_v50,p2);
      let _111_v60;
      _111_v60 = _dafny.MultiSet.fromElements(_110_v59);
      let _112_v61;
      _112_v61 = _dafny.Set.fromElements(p2, p2, p2, p2, p2);
      let _113_v62;
      _113_v62 = _dafny.Seq.of(_dafny.Set.fromElements(_module.__default.fm0(p2, p2, _111_v60, p0, globalState)), _112_v61);
      let _114_v63;
      _114_v63 = _dafny.Seq.of(new BigNumber(-8), _98_v49.f8);
      _109_v58 = (_109_v58).update(_dafny.Seq.contains(_dafny.Seq.of(_112_v61), (_113_v62)[_module.__default.safeIndex(new BigNumber((_114_v63).length), new BigNumber((_113_v62).length))]), (_101_v53)[_module.__default.safeIndex(new BigNumber(952), new BigNumber((_101_v53).length))]);
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_106_v55).length))) {
        let _115_i9 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_115_i9)) && ((_115_i9).isLessThan(new BigNumber((_106_v55).length))))) {
          (_106_v55)[(_115_i9)] = (_112_v61).IsSubsetOf(_112_v61);
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _116_v0;
      _116_v0 = false;
      let _117_v1;
      _117_v1 = _dafny.Map.Empty.slice().updateUnsafe(_116_v0,new BigNumber((_dafny.Seq.UnicodeFromString("gmoircrjr")).length));
      let _118_v2;
      let _init4 = function (_119_i0) {
        return _module.__default.safeModuloInt(_119_i0, new BigNumber(448));
      };
      let _nw21 = Array((_dafny.ONE).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw21.length); _i0_4++) {
        _nw21[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _118_v2 = _nw21;
      let _120_v3;
      let _init5 = function (_121_i1) {
        return false;
      };
      let _nw22 = Array((new BigNumber(12)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw22.length); _i0_5++) {
        _nw22[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _120_v3 = _nw22;
      let _122_v4;
      _122_v4 = _dafny.Seq.UnicodeFromString("pnncmt");
      let _123_globalState;
      let _nw23 = new _module.GlobalState();
      _nw23.__ctor(_117_v1, _118_v2, _120_v3, new BigNumber(287), _122_v4, false, _122_v4, new BigNumber(632));
      _123_globalState = _nw23;
      let _124_v5;
      _124_v5 = new BigNumber(681);
      let _125_v6;
      _125_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_116_v0, _124_v5, _124_v5, _123_globalState),_116_v0);
      let _126_v7;
      _126_v7 = _dafny.Seq.of(_125_v6);
      let _127_v8;
      _127_v8 = _dafny.MultiSet.fromElements(_125_v6, _125_v6, _125_v6);
      _116_v0 = _module.__default.fm0(_module.__default.fm0(_116_v0, _116_v0, _dafny.MultiSet.FromArray(_126_v7), new BigNumber((_dafny.Seq.UnicodeFromString("lcjlcyyhn")).length), _123_globalState), _module.__default.fm0(_116_v0, _116_v0, _dafny.MultiSet.fromElements(_125_v6, _125_v6, _module.__default.fm2(!(_116_v0), _123_globalState)), _124_v5, _123_globalState), (_127_v8).Difference(_dafny.MultiSet.FromArray(_126_v7)), ((_dafny.ZERO).minus(new BigNumber((_122_v4).length))).multipliedBy(_124_v5), _123_globalState);
      let _index22 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
      (_120_v3)[_index22] = true;
      let _index23 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
      (_120_v3)[_index23] = (_124_v5).isLessThanOrEqualTo(_124_v5);
      let _128_v9;
      _128_v9 = _dafny.Seq.of((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _116_v0);
      let _129_v10;
      _129_v10 = _dafny.Map.Empty.slice().updateUnsafe(_116_v0,_118_v2);
      let _130_i2;
      _130_i2 = _dafny.ZERO;
      L1: {
        while ((_128_v9)[_module.__default.safeIndex(new BigNumber(((_129_v10).Merge(_129_v10)).length), new BigNumber((_128_v9).length))]) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_130_i2)) {
              break L1;
            }
            _130_i2 = (_130_i2).plus(_dafny.ONE);
            let _131_v11;
            _131_v11 = _dafny.Seq.of(_124_v5);
            let _132_v12;
            _132_v12 = new _dafny.CodePoint('i'.codePointAt(0));
            let _133_v13;
            _133_v13 = _dafny.Map.Empty.slice().updateUnsafe(_122_v4,_132_v12);
            let _134_v14;
            _134_v14 = _dafny.MultiSet.fromElements((_131_v11)[_module.__default.safeIndex(new BigNumber((_131_v11).length), new BigNumber((_131_v11).length))], _124_v5, (_dafny.ZERO).minus(new BigNumber((_133_v13).length)), _124_v5);
            _module.__default.m0((_dafny.ZERO).minus(_124_v5), _134_v14, ((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]) === ((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]), (_dafny.ZERO).minus(_124_v5), _123_globalState);
            let _135_v15;
            _135_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_122_v4, _122_v4), _module.__default.safeIndex(new BigNumber(-825), new BigNumber((_dafny.Seq.Concat(_122_v4, _122_v4)).length)), _132_v12),(((_128_v9)[_module.__default.safeIndex(_124_v5, new BigNumber((_128_v9).length))]) ? (!(_116_v0)) : ((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))])));
            let _136_v17;
            _136_v17 = _dafny.MultiSet.fromElements(_122_v4, _dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("npfn"));
            let _rhs11 = (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))];
            let _rhs12 = (_dafny.ZERO).minus((_124_v5).minus(_124_v5));
            let _rhs13 = function () {
              let _coll4 = new _dafny.Map();
              for (const _compr_4 of (_136_v17).Elements) {
                let _137_v16 = _compr_4;
                if ((_136_v17).contains(_137_v16)) {
                  _coll4.push([_137_v16,_116_v0]);
                }
              }
              return _coll4;
            }();
            let _lhs13 = _123_globalState;
            _116_v0 = _rhs11;
            _lhs13.f7 = _rhs12;
            _135_v15 = _rhs13;
            if (!(!((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]) || (_116_v0)) || ((_134_v14).contains(_124_v5))) {
              let _index24 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_118_v2).length));
              (_118_v2)[_index24] = _124_v5;
              let _index25 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_118_v2).length));
              (_118_v2)[_index25] = _124_v5;
              let _index26 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_118_v2).length));
              (_118_v2)[_index26] = _124_v5;
              (_123_globalState).f7 = ((_118_v2)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_118_v2).length))]).minus(_module.__default.safeModuloInt(_module.__default.fm3(_123_globalState), (_dafny.ZERO).minus(_124_v5)));
              let _138_v18;
              let _init6 = ((_139_v0) => function (_140_i3) {
                return _dafny.Seq.of(_139_v0);
              })(_116_v0);
              let _nw24 = Array((new BigNumber(7)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw24.length); _i0_6++) {
                _nw24[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _138_v18 = _nw24;
              let _index27 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v18).length));
              (_138_v18)[_index27] = _dafny.Seq.Concat(_128_v9, _128_v9);
              let _index28 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v18).length));
              (_138_v18)[_index28] = _128_v9;
              _116_v0 = ((_138_v18)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v18).length))])[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_138_v18)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v18).length))])).cardinality()), new BigNumber(((_138_v18)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v18).length))]).length))];
            } else {
              let _141_v19;
              let _nw25 = new _module.C0();
              _nw25.__ctor(_124_v5);
              _141_v19 = _nw25;
              let _142_v20;
              let _143_v21;
              let _144_v22;
              let _145_v23;
              let _out0;
              let _out1;
              let _out2;
              let _out3;
              let _outcollector0 = (_141_v19).m1(_123_globalState);
              _out0 = _outcollector0[0];
              _out1 = _outcollector0[1];
              _out2 = _outcollector0[2];
              _out3 = _outcollector0[3];
              _142_v20 = _out0;
              _143_v21 = _out1;
              _144_v22 = _out2;
              _145_v23 = _out3;
              _131_v11 = _dafny.Seq.Concat(_131_v11, _131_v11);
              let _146_v24;
              _146_v24 = _dafny.Map.Empty.slice().updateUnsafe(_141_v19,_141_v19);
              let _147_v25;
              _147_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(587),_146_v24);
              let _148_v26;
              _148_v26 = _dafny.MultiSet.fromElements(_141_v19, _141_v19);
              _module.__default.m0(new BigNumber((_dafny.Seq.Concat(_122_v4, _122_v4)).length), _134_v14, !(new BigNumber((_dafny.Seq.update(_122_v4, _module.__default.safeIndex(_141_v19.f8, new BigNumber((_122_v4).length)), new _dafny.CodePoint('c'.codePointAt(0)))).length)).isEqualTo(new BigNumber(((((_147_v25).contains(_141_v19.f8)) ? ((_147_v25).get(_141_v19.f8)) : (_146_v24))).length)), new BigNumber((_148_v26).cardinality()), _123_globalState);
              let _149_v27;
              _149_v27 = _module.D2.create_DC5(_128_v9);
              let _150_v28;
              _150_v28 = _dafny.Set.fromElements(_143_v21, !(_143_v21) || (false), !(_124_v5).isEqualTo((((_117_v1).contains(_116_v0)) ? ((_117_v1).get(_116_v0)) : (new BigNumber(((_149_v27).dtor_cf9).length)))));
              _150_v28 = _150_v28;
            }
            _module.__default.m0(_124_v5, _134_v14, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _124_v5, _123_globalState);
          }
        }
      }
      let _151_v29;
      let _nw26 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _151_v29 = _nw26;
      let _152_v30;
      _152_v30 = _dafny.Set.fromElements(_151_v29, _151_v29, _151_v29, _151_v29, _151_v29);
      let _153_v31;
      _153_v31 = _dafny.MultiSet.fromElements(_124_v5);
      let _154_v32;
      _154_v32 = _module.D0.create_DC3(_117_v1, _116_v0, _124_v5);
      _module.__default.m0(new BigNumber((_152_v30).length), _153_v31, (_154_v32).dtor_cf6, _124_v5, _123_globalState);
      let _155_v33;
      _155_v33 = _dafny.Map.Empty.slice().updateUnsafe(_124_v5,_118_v2);
      _155_v33 = (_155_v33).update(_124_v5, _118_v2);
      let _156_v34;
      _156_v34 = _module.D0.create_DC1(_124_v5, (_124_v5).minus(_124_v5), _124_v5);
      _156_v34 = _module.D0.create_DC1(new BigNumber(322), _124_v5, _124_v5);
      let _157_v35;
      _157_v35 = _module.D0.create_DC2(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-986)), ((_158_v5) => function (_159_i4) {
  return _158_v5;
})(_124_v5))).length));
      let _160_v36;
      _160_v36 = _dafny.MultiSet.fromElements(_157_v35, _module.D0.create_DC2(_124_v5));
      _160_v36 = (_160_v36).update(_157_v35, _module.__default.abs(_124_v5));
      let _161_v37;
      _161_v37 = _dafny.Map.Empty.slice().updateUnsafe(_116_v0,_116_v0);
      let _rhs14 = ((false) ? (new BigNumber(((_161_v37).Merge(_161_v37)).length)) : (_124_v5));
      let _rhs15 = _124_v5;
      let _lhs14 = _123_globalState;
      _124_v5 = _rhs14;
      _lhs14.f7 = _rhs15;
      let _162_v38;
      _162_v38 = _dafny.Map.Empty.slice().updateUnsafe(_118_v2,(_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]);
      let _163_v39;
      _163_v39 = new _dafny.CodePoint('c'.codePointAt(0));
      if ((((_162_v38).contains(_118_v2)) ? ((_162_v38).get(_118_v2)) : (_module.__default.fm0(!(_116_v0), _module.__default.fm0(false, false, _dafny.MultiSet.fromElements((_125_v6).update(_163_v39, _116_v0)), _124_v5, _123_globalState), (_dafny.MultiSet.FromArray(_126_v7)).update(_125_v6, _module.__default.abs(_124_v5)), _124_v5, _123_globalState)))) {
        let _164_v40;
        _164_v40 = _dafny.Map.Empty.slice().updateUnsafe(_116_v0,_120_v3);
        _164_v40 = (_164_v40).update(_116_v0, _120_v3);
        let _165_v41;
        _165_v41 = _module.D3.create_DC7(_161_v37);
        _116_v0 = (((_165_v41).dtor_cf11).Merge(_161_v37)).equals((_161_v37).Merge(_module.__default.fm4(!(_116_v0), _124_v5, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _123_globalState)));
        let _166_v42;
        _166_v42 = _dafny.MultiSet.fromElements((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _116_v0, !(!(!(_116_v0))) || ((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]), _116_v0, _dafny.Seq.IsProperPrefixOf(_122_v4, _122_v4));
        let _167_v43;
        _167_v43 = _dafny.Set.fromElements(!(_116_v0));
        let _168_v44;
        _168_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(604)), ((_169_v5) => function (_170_i5) {
          return _169_v5;
        })(_124_v5))).length),!(_116_v0));
        (_123_globalState).f7 = (((_166_v42).contains(_module.__default.fm0(_116_v0, _116_v0, _127_v8, new BigNumber((_167_v43).length), _123_globalState))) ? ((_166_v42).get(_module.__default.fm0(_116_v0, _116_v0, _127_v8, new BigNumber((_167_v43).length), _123_globalState))) : (_module.__default.safeDivisionInt(new BigNumber((_module.__default.fm5(new BigNumber(17), _116_v0, (((_168_v44).contains(new BigNumber((_153_v31).cardinality()))) ? ((_168_v44).get(new BigNumber((_153_v31).cardinality()))) : (true)), _123_globalState)).cardinality()), new BigNumber(136))));
        let _171_v46;
        _171_v46 = _dafny.Map.Empty.slice().updateUnsafe(_124_v5,_module.D0.create_DC0(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_172_i7) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("hwbs")).length);
}))));
        let _173_v47;
        _173_v47 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_124_v5,(_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))])).length));
        let _174_v48;
        _174_v48 = _dafny.Seq.of(_173_v47);
        let _175_v49;
        _175_v49 = _module.D0.create_DC0(_174_v48);
        _module.__default.m0(new BigNumber(((function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(954), new BigNumber(623))) {
            let _176_v45 = _compr_5;
            if (((new BigNumber(954)).isLessThanOrEqualTo(_176_v45)) && ((_176_v45).isLessThan(new BigNumber(623)))) {
              _coll5.push([(_176_v45).plus(new BigNumber((_dafny.Seq.of(_124_v5)).length)),_module.D0.create_DC0(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-263)), ((_177_v5) => function (_178_i6) {
  return _177_v5;
})(_124_v5)), _dafny.Seq.of(_124_v5, new BigNumber((_122_v4).length))))]);
            }
          }
          return _coll5;
        }()).Merge((_171_v46).update(new BigNumber((_122_v4).length), _175_v49))).length), _153_v31, _116_v0, _124_v5, _123_globalState);
        _124_v5 = _124_v5;
      } else {
        let _179_v50;
        let _nw27 = new _module.C0();
        _nw27.__ctor(_module.__default.safeModuloInt(_124_v5, new BigNumber(-119)));
        _179_v50 = _nw27;
        let _180_v51;
        _180_v51 = _dafny.Set.fromElements(_161_v37);
        let _181_v52;
        _181_v52 = _180_v51;
        let _182_v53;
        _182_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-303),_181_v52);
        let _183_v54;
        _183_v54 = _dafny.Seq.of(_124_v5, new BigNumber(801), _179_v50.f8, _179_v50.f8, new BigNumber(333));
        _182_v53 = (_182_v53).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_183_v54,_122_v4)).length), _181_v52);
        (_123_globalState).f7 = (_156_v34).dtor_cf2;
        (_123_globalState).f7 = _179_v50.f8;
        let _184_v55;
        let _nw28 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
        _184_v55 = _nw28;
        let _185_v56;
        _185_v56 = _dafny.Map.Empty.slice().updateUnsafe(_179_v50,_184_v55);
        let _rhs16 = _157_v35;
        let _rhs17 = _154_v32;
        let _rhs18 = _179_v50.f8;
        let _rhs19 = ((_116_v0) ? ((((_185_v56).contains(_179_v50)) ? ((_185_v56).get(_179_v50)) : (_184_v55))) : (((_116_v0) ? (_184_v55) : (_184_v55))));
        let _lhs15 = _179_v50;
        _157_v35 = _rhs16;
        _154_v32 = _rhs17;
        _lhs15.f8 = _rhs18;
        _184_v55 = _rhs19;
      }
      _124_v5 = _124_v5;
      let _186_v57;
      let _nw29 = new _module.C0();
      _nw29.__ctor(new BigNumber(530));
      _186_v57 = _nw29;
      _118_v2 = _118_v2;
      let _hi1 = _124_v5;
      for (let _187_i8 = _186_v57.f8; _187_i8.isLessThan(_hi1); _187_i8 = _187_i8.plus(_dafny.ONE)) {
        let _188_v58;
        _188_v58 = _dafny.Map.Empty.slice().updateUnsafe(_186_v57,false);
        _188_v58 = (_188_v58).update(_186_v57, !((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]));
        let _189_v59;
        let _init7 = ((_190_v4, _191_v5, _192_v39) => function (_193_i9) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.update(_190_v4, _module.__default.safeIndex(_191_v5, new BigNumber((_190_v4).length)), _192_v39)), _dafny.Seq.of(_190_v4, _dafny.Seq.UnicodeFromString("fkc")));
        })(_122_v4, _124_v5, _163_v39);
        let _nw30 = Array((new BigNumber(13)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw30.length); _i0_7++) {
          _nw30[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _189_v59 = _nw30;
        let _194_v60;
        _194_v60 = _dafny.Seq.of(_122_v4, (_module.__default.fm6(_187_i8, _186_v57.f8, _116_v0, new _dafny.CodePoint('f'.codePointAt(0)), _123_globalState)).dtor_cf12, _122_v4, _122_v4);
        let _index29 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_189_v59).length));
        (_189_v59)[_index29] = (((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]) ? (_194_v60) : (_194_v60));
        let _index30 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_189_v59).length));
        (_189_v59)[_index30] = _194_v60;
        let _index31 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
        (_120_v3)[_index31] = _116_v0;
        let _195_v61;
        _195_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_118_v2,_163_v39),_116_v0);
        let _196_v62;
        _196_v62 = _dafny.Map.Empty.slice().updateUnsafe(_118_v2,_163_v39);
        _116_v0 = !((_128_v9)[_module.__default.safeIndex(_186_v57.f8, new BigNumber((_128_v9).length))]) || ((((_195_v61).contains(_196_v62)) ? ((_195_v61).get(_196_v62)) : ((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))])));
      }
      let _hi2 = _186_v57.f8;
      for (let _197_i10 = _186_v57.f8; _197_i10.isLessThan(_hi2); _197_i10 = _197_i10.plus(_dafny.ONE)) {
        let _198_v63;
        _198_v63 = _dafny.Seq.of(_161_v37);
        let _199_v64;
        _199_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_198_v63).length),_116_v0);
        let _200_v65;
        _200_v65 = _dafny.Map.Empty.slice().updateUnsafe(_186_v57.f8,_124_v5);
        _199_v64 = (_199_v64).update(_module.__default.safeDivisionInt(new BigNumber((_200_v65).length), _197_i10), _116_v0);
        let _index32 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
        (_120_v3)[_index32] = _module.__default.fm0((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _127_v8, _186_v57.f8, _123_globalState);
        let _index33 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
        (_120_v3)[_index33] = (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))];
        _157_v35 = _157_v35;
      }
      let _201_v66;
      let _nw31 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _201_v66 = _nw31;
      let _202_v67;
      _202_v67 = _dafny.Seq.of(_124_v5);
      let _index34 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_201_v66).length));
      (_201_v66)[_index34] = _202_v67;
      let _index35 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_201_v66).length));
      (_201_v66)[_index35] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), ((_203_v57) => function (_204_i11) {
        return _203_v57.f8;
      })(_186_v57)), _202_v67);
      let _hi3 = (_124_v5).plus(new BigNumber(-16));
      for (let _205_i12 = (_dafny.ZERO).minus(_124_v5); _205_i12.isLessThan(_hi3); _205_i12 = _205_i12.plus(_dafny.ONE)) {
        if (!(!_dafny.Seq.contains(_122_v4, _163_v39))) {
          let _206_v68;
          _206_v68 = _dafny.Map.Empty.slice().updateUnsafe(_116_v0,_186_v57);
          let _207_v69;
          _207_v69 = _dafny.Map.Empty.slice().updateUnsafe((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))],_206_v68);
          _207_v69 = (_207_v69).update(false, (_206_v68).update((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _186_v57));
          let _208_v70;
          _208_v70 = _dafny.MultiSet.fromElements(_116_v0, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _116_v0);
          let _index36 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
          let _rhs20 = !((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]);
          let _rhs21 = _module.__default.safeModuloInt(new BigNumber(454), _186_v57.f8);
          let _rhs22 = new BigNumber(-651);
          let _rhs23 = ((_186_v57.f8).minus(_124_v5)).minus(_186_v57.f8);
          let _rhs24 = ((true) ? ((((_161_v37).contains(_116_v0)) ? ((_161_v37).get(_116_v0)) : (_module.__default.fm0(_116_v0, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _127_v8, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(109)), ((_209_v39) => function (_210_i13) {
            return _209_v39;
          })(_163_v39))).length), _123_globalState)))) : ((_208_v70).IsSubsetOf(_208_v70)));
          let _lhs16 = _186_v57;
          let _lhs17 = _186_v57;
          let _lhs18 = _186_v57;
          let _lhs19 = _120_v3;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
          _116_v0 = _rhs20;
          _lhs16.f8 = _rhs21;
          _lhs17.f8 = _rhs22;
          _lhs18.f8 = _rhs23;
          _lhs19[_lhs20] = _rhs24;
          let _index37 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
          (_120_v3)[_index37] = (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))];
          let _211_v71;
          _211_v71 = _module.D2.create_DC5(_dafny.Seq.of((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]));
          let _212_v72;
          _212_v72 = _dafny.Map.Empty.slice().updateUnsafe(_163_v39,_211_v71);
          _212_v72 = (_212_v72).update(_163_v39, _211_v71);
          _124_v5 = new BigNumber(-669);
        } else {
          let _213_v74;
          _213_v74 = _module.D2.create_DC5(_128_v9);
          let _214_v75;
          _214_v75 = _dafny.Map.Empty.slice().updateUnsafe(_122_v4,(_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]);
          let _215_v76;
          _215_v76 = _module.D4.create_DC9(_214_v75);
          let _216_v78;
          _216_v78 = _dafny.Set.fromElements(_122_v4);
          let _217_v79;
          _217_v79 = _dafny.Seq.of(function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_216_v78).Elements) {
              let _218_v77 = _compr_6;
              if ((_216_v78).contains(_218_v77)) {
                _coll6.push([_218_v77,(_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]]);
              }
            }
            return _coll6;
          }(), _214_v75);
          let _219_v80;
          let _nw32 = Array((new BigNumber(17)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = (function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of (_module.__default.fm7(_116_v0, _213_v74, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _123_globalState)).Keys.Elements) {
              let _220_v73 = _compr_7;
              if ((_module.__default.fm7(_116_v0, _213_v74, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _123_globalState)).contains(_220_v73)) {
                _coll7.push([_220_v73,(_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))]]);
              }
            }
            return _coll7;
          }()).Merge(_214_v75);
          _nw32[(_dafny.ONE).toNumber()] = _214_v75;
          _nw32[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_122_v4,false);
          _nw32[(new BigNumber(3)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(4)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(5)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(6)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(7)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(8)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(9)).toNumber()] = (_214_v75).Merge(_214_v75);
          _nw32[(new BigNumber(10)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(11)).toNumber()] = (_214_v75).Merge(_214_v75);
          _nw32[(new BigNumber(12)).toNumber()] = (_215_v76).dtor_cf14;
          _nw32[(new BigNumber(13)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(14)).toNumber()] = (_217_v79)[_module.__default.safeIndex(_124_v5, new BigNumber((_217_v79).length))];
          _nw32[(new BigNumber(15)).toNumber()] = _214_v75;
          _nw32[(new BigNumber(16)).toNumber()] = _214_v75;
          _219_v80 = _nw32;
          let _index38 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_219_v80).length));
          (_219_v80)[_index38] = _module.__default.fm8(_module.__default.fm0(false, false, _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_163_v39,_116_v0), _125_v6), new BigNumber(10), _123_globalState), _123_globalState);
          let _index39 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_219_v80).length));
          (_219_v80)[_index39] = (_214_v75).Merge(_214_v75);
          _122_v4 = _dafny.Seq.Concat(_122_v4, _122_v4);
          let _221_v81;
          _221_v81 = _module.D3.create_DC7(_161_v37);
          let _222_v82;
          _222_v82 = _module.D3.create_DC8(_122_v4, _202_v67);
          let _223_v83;
          _223_v83 = _dafny.MultiSet.fromElements(_222_v82, _222_v82);
          _221_v81 = _module.__default.fm9(_186_v57.f8, new BigNumber((_223_v83).cardinality()), _123_globalState);
          (_123_globalState).f7 = (_dafny.ZERO).minus(_186_v57.f8);
          (_186_v57).f8 = _186_v57.f8;
        }
        (_186_v57).f8 = _186_v57.f8;
        let _224_v84;
        _224_v84 = _dafny.Map.Empty.slice().updateUnsafe(_186_v57,!(_116_v0) || (_116_v0));
        _224_v84 = (_224_v84).update(_186_v57, _116_v0);
        let _225_v85;
        let _init8 = ((_226_v37) => function (_227_i14) {
          return _dafny.Set.fromElements(_226_v37);
        })(_161_v37);
        let _nw33 = Array((new BigNumber(20)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw33.length); _i0_8++) {
          _nw33[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _225_v85 = _nw33;
        let _228_v86;
        _228_v86 = _module.D5.create_DC11(_163_v39);
        let _index40 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_151_v29).length));
        (_151_v29)[_index40] = (_228_v86).dtor_cf17;
        let _229_v87;
        _229_v87 = _dafny.MultiSet.fromElements(_module.__default.fm0(_116_v0, (_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))], _127_v8, new BigNumber((_202_v67).length), _123_globalState));
        let _index41 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_151_v29).length));
        let _index42 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
        let _rhs25 = (_124_v5).multipliedBy((((_229_v87).contains((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))])) ? ((_229_v87).get((_120_v3)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length))])) : (_205_i12)));
        let _rhs26 = _225_v85;
        let _rhs27 = new _dafny.CodePoint('r'.codePointAt(0));
        let _rhs28 = _186_v57.f8;
        let _rhs29 = true;
        let _lhs21 = _186_v57;
        let _lhs22 = _151_v29;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_151_v29).length));
        let _lhs24 = _123_globalState;
        let _lhs25 = _120_v3;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_120_v3).length));
        _lhs21.f8 = _rhs25;
        _225_v85 = _rhs26;
        _lhs22[_lhs23] = _rhs27;
        _lhs24.f7 = _rhs28;
        _lhs25[_lhs26] = _rhs29;
      }
      process.stdout.write(_dafny.toString(_116_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_117_v1).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_122_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f0).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_123_globalState).f2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_123_globalState.f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_123_globalState).f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_123_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_124_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_126_v7, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v8).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_128_v9, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_129_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_130_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v29)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_152_v30).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v31).equals(_dafny.MultiSet.fromElements(new BigNumber(681)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v32).dtor_cf5).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v32).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v32).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_155_v33).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v34).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v34).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v34).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_v35).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v36).equals(_dafny.MultiSet.fromElements(_module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(986)), _module.D0.create_DC2(new BigNumber(681))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v37).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_162_v38).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_163_v39));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_186_v57.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_201_v66)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(530), new BigNumber(681)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_202_v67, _dafny.Seq.of(new BigNumber(681)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf5, cf6, cf7) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf10) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC8(cf12, cf13) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D3(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + this.cf12.toVerbatimString(true) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.Seq.UnicodeFromString(""), _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf15, cf16) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D4(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(false, null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf18, cf19, cf20) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC13(cf21, cf22, cf23, cf24) {
      let $dt = new D5(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC14(cf25, cf26, cf27) {
      let $dt = new D5(2);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D5(3);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D5(4);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get is_DC15() { return this.$tag === 4; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 4) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf30) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf31, cf32) {
      let $dt = new D6(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19() {
      let $dt = new D6(2);
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D6(3);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19";
      } else if (this.$tag === 3) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31 && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f4 = _dafny.Seq.UnicodeFromString("");
      this.f7 = _dafny.ZERO;
      this._f0 = _dafny.Map.Empty;
      this._f1 = [];
      this._f2 = [];
      this._f3 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f8) {
      let _this = this;
      (_this).f8 = f8;
      return;
    }
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _230_v0;
      _230_v0 = _dafny.Seq.UnicodeFromString("wdjxni");
      let _231_v1;
      _231_v1 = true;
      let _232_v2;
      _232_v2 = _dafny.Map.Empty.slice().updateUnsafe(_231_v1,_this.f8);
      let _233_v3;
      let _nw34 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _233_v3 = _nw34;
      let _234_v4;
      _234_v4 = _dafny.Map.Empty.slice().updateUnsafe(_233_v3,_this.f8);
      let _235_v5;
      _235_v5 = _dafny.Seq.of(new BigNumber((_232_v2).length), new BigNumber((_234_v4).length));
      let _236_v6;
      let _nw35 = Array((new BigNumber(4)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf(_230_v0, _230_v0);
      _nw35[(_dafny.ONE).toNumber()] = !_dafny.Seq.contains((_module.D0.create_DC0(_dafny.Seq.of(_235_v5))).dtor_cf0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(117)), function (_237_i0) {
        return _this.f8;
      }));
      _nw35[(new BigNumber(2)).toNumber()] = _231_v1;
      _nw35[(new BigNumber(3)).toNumber()] = _231_v1;
      _236_v6 = _nw35;
      let _238_v7;
      _238_v7 = _dafny.Map.Empty.slice().updateUnsafe(!(_231_v1),_231_v1);
      let _239_v8;
      _239_v8 = _dafny.Set.fromElements(_238_v7, ((_238_v7).update(_231_v1, false)).update(_231_v1, true), _238_v7);
      let _240_v9;
      _240_v9 = new _dafny.CodePoint('f'.codePointAt(0));
      let _241_v11;
      _241_v11 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_231_v1), _238_v7);
      let _rhs30 = new BigNumber((function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_240_v9,!(_231_v1)),_231_v1)).Keys.Elements) {
          let _242_v10 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_240_v9,!(_231_v1)),_231_v1)).contains(_242_v10)) {
            _coll8.add(_242_v10);
          }
        }
        return _coll8;
      }()).length);
      let _rhs31 = _236_v6;
      let _rhs32 = _module.__default.safeDivisionInt((_this.f8).multipliedBy(_this.f8), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_231_v1)).length)));
      let _rhs33 = (_this.f8).plus(_this.f8);
      let _rhs34 = (_241_v11);
      let _lhs27 = globalState;
      let _lhs28 = globalState;
      let _lhs29 = globalState;
      _lhs27.f7 = _rhs30;
      _236_v6 = _rhs31;
      _lhs28.f7 = _rhs32;
      _lhs29.f7 = _rhs33;
      _239_v8 = _rhs34;
      let _243_v12;
      _243_v12 = _dafny.Seq.of(_231_v1);
      let _244_v13;
      _244_v13 = _dafny.Seq.of(_243_v12);
      _243_v12 = _dafny.Seq.Concat(_243_v12, (_244_v13)[_module.__default.safeIndex(_this.f8, new BigNumber((_244_v13).length))]);
      (globalState).f7 = _this.f8;
      (globalState).f7 = ((_dafny.ZERO).minus(_this.f8)).multipliedBy(new BigNumber(616));
      (_this).f8 = _this.f8;
      let _245_v14;
      _245_v14 = _dafny.MultiSet.fromElements(_this.f8);
      _module.__default.m0(_this.f8, _245_v14, _231_v1, _this.f8, globalState);
      let _246_v15;
      _246_v15 = _dafny.Map.Empty.slice().updateUnsafe(_240_v9,_231_v1);
      let _247_v16;
      _247_v16 = _dafny.MultiSet.fromElements(_246_v15);
      r0 = _module.__default.fm0(_module.__default.fm0(true, _231_v1, _247_v16, new BigNumber((_239_v8).length), globalState), _231_v1, ((_247_v16).update(function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), ((_248_v9) => function (_249_i1) {
          return _248_v9;
        })(_240_v9))).Elements) {
          let _250_v17 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), ((_251_v9) => function (_249_i1) {
            return _251_v9;
          })(_240_v9)), _250_v17)) {
            _coll9.push([_250_v17,!(_231_v1)]);
          }
        }
        return _coll9;
      }(), _module.__default.abs(_this.f8))).Union(_247_v16), _this.f8, globalState);
      r1 = _231_v1;
      let _252_v18;
      _252_v18 = _dafny.Set.fromElements(_231_v1, true);
      r2 = _module.__default.safeDivisionInt(new BigNumber((_252_v18).length), ((_dafny.ZERO).minus(_this.f8)).minus(_this.f8));
      let _253_v19;
      _253_v19 = _module.D0.create_DC3((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(!(_231_v1), false, _231_v1, !(_231_v1))).length))).update(_231_v1, _this.f8), _231_v1, _this.f8);
      let _pat_let_tv3 = _240_v9;
      let _pat_let_tv4 = _240_v9;
      let _pat_let_tv5 = _230_v0;
      let _pat_let_tv6 = _230_v0;
      r3 = function (_source2) {
        if (_source2.is_DC1) {
          let _254___mcc_h0 = (_source2).cf1;
          let _255___mcc_h1 = (_source2).cf2;
          let _256___mcc_h2 = (_source2).cf3;
          let _257_cf3 = _256___mcc_h2;
          let _258_cf2 = _255___mcc_h1;
          let _259_cf1 = _254___mcc_h0;
          return _pat_let_tv3;
        } else if (_source2.is_DC2) {
          let _260___mcc_h3 = (_source2).cf4;
          let _261_cf4 = _260___mcc_h3;
          return new _dafny.CodePoint('d'.codePointAt(0));
        } else if (_source2.is_DC3) {
          let _262___mcc_h4 = (_source2).cf5;
          let _263___mcc_h5 = (_source2).cf6;
          let _264___mcc_h6 = (_source2).cf7;
          let _265_cf7 = _264___mcc_h6;
          let _266_cf6 = _263___mcc_h5;
          let _267_cf5 = _262___mcc_h4;
          return _pat_let_tv4;
        } else {
          let _268___mcc_h7 = (_source2).cf0;
          let _269_cf0 = _268___mcc_h7;
          return (_pat_let_tv5)[_module.__default.safeIndex(_this.f8, new BigNumber((_pat_let_tv6).length))];
        }
      }(_253_v19);
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
