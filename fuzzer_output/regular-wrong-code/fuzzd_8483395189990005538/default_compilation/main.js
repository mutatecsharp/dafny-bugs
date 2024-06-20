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
    static fm0(p0, p1, globalState) {
      return (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).length)).isLessThanOrEqualTo(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(555), new BigNumber(992))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(555)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(992)))) {
            _coll0.push([(_0_v0).multipliedBy(new BigNumber(-542)),_module.D2.create_DC6(_dafny.Seq.of(new BigNumber(123), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(623)), function (_1_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_2_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}),true)).length), new BigNumber((_dafny.Seq.of(false)).length)))]);
          }
        }
        return _coll0;
      }()).length));
    };
    static fm1(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(55)),true)).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of((_module.D2.create_DC6(_dafny.Seq.of(new BigNumber(523), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).cardinality())))).dtor_cf11, _dafny.Seq.of(new BigNumber(156), new BigNumber(-520)))).Elements) {
          let _3_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of((_module.D2.create_DC6(_dafny.Seq.of(new BigNumber(523), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).cardinality())))).dtor_cf11, _dafny.Seq.of(new BigNumber(156), new BigNumber(-520))), _3_v0)) {
            _coll1.push([_3_v0,!(true)]);
          }
        }
        return _coll1;
      }())).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), function (_4_i0) {
          return new BigNumber(742);
        }),new BigNumber(225))).Keys.Elements) {
          let _5_v1 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), function (_4_i0) {
            return new BigNumber(742);
          }),new BigNumber(225))).contains(_5_v1)) {
            _coll2.push([_5_v1,true]);
          }
        }
        return _coll2;
      }());
    };
    static fm2(p0, p1, globalState) {
      if ((!(!(false))) || (true)) {
        return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),true)).Keys.Elements) {
            let _6_v0 = _compr_3;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),true)).contains(_6_v0)) {
              _coll3.add(_6_v0);
            }
          }
          return _coll3;
        }())).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0))))));
      } else {
        return _dafny.MultiSet.fromElements(function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), function (_7_i0) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }))).Elements) {
            let _8_v1 = _compr_4;
            if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), function (_7_i0) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            }))).contains(_8_v1)) {
              _coll4.add(_8_v1);
            }
          }
          return _coll4;
        }(), _dafny.Set.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0))), function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.Seq.UnicodeFromString("iasqoksj"))).Keys.Elements) {
            let _9_v2 = _compr_5;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.Seq.UnicodeFromString("iasqoksj"))).contains(_9_v2)) {
              _coll5.add(_9_v2);
            }
          }
          return _coll5;
        }());
      }
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_10_i0) {
        return function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new _dafny.CodePoint('q'.codePointAt(0)))).Keys.Elements) {
            let _11_v0 = _compr_6;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new _dafny.CodePoint('q'.codePointAt(0)))).contains(_11_v0)) {
              _coll6.add(_11_v0);
            }
          }
          return _coll6;
        }();
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(872)), function (_12_i1) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)));
      }));
    };
    static fm11(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm13(p0, p1, p2, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("anbp"), _dafny.Seq.UnicodeFromString("qmbavh")), _dafny.Seq.UnicodeFromString("e"))).length);
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(((!(false)) ? (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false, true)).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), function (_13_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-169)))).length));
      }))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, !(!(true)))).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-274))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(525),new BigNumber(504))).length), new BigNumber(-521)), _dafny.Seq.of(new BigNumber(174), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, false)).length)),new _dafny.CodePoint('p'.codePointAt(0)))).length))));
    };
    static fm17(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("jogx");
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sxblvklw")).length),new BigNumber(135)));
    };
    static fm19(globalState) {
      return _module.D1.create_DC2((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hvymmqds")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),!(true))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ygsxci")).length))).cardinality()),new BigNumber((_dafny.Seq.of(true, false)).length))));
    };
    static fm20(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(!(false)))));
    };
    static fm21(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), function (_14_i0) {
        return (function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of (_dafny.Seq.of(new BigNumber(-371))).Elements) {
            let _15_v0 = _compr_7;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-371)), _15_v0)) {
              _coll7.add((_15_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(792))).cardinality())),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(504))).length))).length)));
            }
          }
          return _coll7;
        }()).Union(_dafny.Set.fromElements(new BigNumber(101), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bkuhorvqx")).length))).Elements) {
            let _16_v1 = _compr_8;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bkuhorvqx")).length))).contains(_16_v1)) {
              _coll8.push([(_16_v1).minus(new BigNumber(-619)),_module.D9.create_DC23(_dafny.MultiSet.fromElements(true))]);
            }
          }
          return _coll8;
        }()).length))).length)));
      });
    };
    static fm22(p0, p1, globalState) {
      if ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-459)), function (_17_i0) {
        return new BigNumber((_dafny.Seq.of(true)).length);
      }), _dafny.Seq.of(new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(786), new BigNumber(626))) {
          let _18_v0 = _compr_9;
          if (((new BigNumber(786)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(626)))) {
            _coll9.add(_module.__default.safeDivisionInt(_18_v0, new BigNumber((_dafny.Seq.UnicodeFromString("olteeivdc")).length)));
          }
        }
        return _coll9;
      }()).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-830)), function (_19_i1) {
        return new BigNumber(386);
      })).length)))).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)), _dafny.Seq.of(new BigNumber(885), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("mxe"))).length), new BigNumber(-13)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(12))).length), new BigNumber(170), new BigNumber((_dafny.Seq.UnicodeFromString("tjay")).length)), _dafny.Seq.of(new BigNumber(890))))) {
        if (true) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }
      } else if (true) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }
    };
    static fm23(p0, p1, globalState) {
      if (true) {
        return _module.D2.create_DC7(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0))), new BigNumber((_dafny.Seq.UnicodeFromString("eurek")).length), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).length), true);
      } else {
        return _module.D2.create_DC7(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0))), new BigNumber((_dafny.Seq.of(!(false))).length), new BigNumber(798), !(false));
      }
    };
    static fm24(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(102))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(47))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(686))));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(false, (_module.D9.create_DC25(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length))).length), true)).dtor_cf41, false, true)).Difference(_dafny.Set.fromElements(true));
    };
    static fm26(p0, p1, p2, globalState) {
      return ((function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(782), new BigNumber(-958))) {
          let _20_v0 = _compr_10;
          if (((new BigNumber(782)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(-958)))) {
            _coll10.add((_20_v0).minus(new BigNumber(86)));
          }
        }
        return _coll10;
      }()).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(411), new BigNumber(918), new BigNumber(498), new BigNumber(372), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), function (_21_i0) {
        return new BigNumber(75);
      }))).cardinality())))).cardinality()), new BigNumber(478)))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dslxbs")).length)));
    };
    static fm28(p0, p1, p2, globalState) {
      return _dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)));
    };
    static fm31(globalState) {
      return _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.of(true, true)).length)).multipliedBy(new BigNumber(742)), new BigNumber(652));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jjjqasxw"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-236)), function (_22_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("idsrwnse")));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D2.create_DC7(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))), (_dafny.ZERO).minus(new BigNumber(-221)), new BigNumber((_dafny.Seq.of(new BigNumber(625))).length), false);
      if (_source0.is_DC7) {
        let _23___mcc_h0 = (_source0).cf12;
        let _24___mcc_h1 = (_source0).cf13;
        let _25___mcc_h2 = (_source0).cf14;
        let _26___mcc_h3 = (_source0).cf15;
        let _27_cf15 = _26___mcc_h3;
        let _28_cf14 = _25___mcc_h2;
        let _29_cf13 = _24___mcc_h1;
        let _30_cf12 = _23___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_29_cf13,_27_cf15)).Merge(function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(892))) {
            let _31_v0 = _compr_11;
            if (((new BigNumber(109)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(892)))) {
              _coll11.push([_module.__default.safeDivisionInt(_31_v0, _28_cf14),_27_cf15]);
            }
          }
          return _coll11;
        }());
      } else if (_source0.is_DC8) {
        let _32___mcc_h4 = (_source0).cf16;
        let _33___mcc_h5 = (_source0).cf17;
        let _34___mcc_h6 = (_source0).cf18;
        let _35_cf18 = _34___mcc_h6;
        let _36_cf17 = _33___mcc_h5;
        let _37_cf16 = _32___mcc_h4;
        return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_36_cf17,new BigNumber(741))).length)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(142),_36_cf17));
      } else {
        let _38___mcc_h7 = (_source0).cf11;
        let _39_cf11 = _38___mcc_h7;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(138),true);
      }
    };
    static fm34(globalState) {
      let _source1 = ((true) ? (_module.D11.create_DC30(true)) : (_module.D11.create_DC30(true)));
      if (_source1.is_DC28) {
        let _40___mcc_h0 = (_source1).cf44;
        let _41_cf44 = _40___mcc_h0;
        return new _dafny.CodePoint('y'.codePointAt(0));
      } else if (_source1.is_DC29) {
        let _42___mcc_h1 = (_source1).cf45;
        let _43___mcc_h2 = (_source1).cf46;
        let _44___mcc_h3 = (_source1).cf47;
        let _45___mcc_h4 = (_source1).cf48;
        let _46_cf48 = _45___mcc_h4;
        let _47_cf47 = _44___mcc_h3;
        let _48_cf46 = _43___mcc_h2;
        let _49_cf45 = _42___mcc_h1;
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else if (_source1.is_DC30) {
        let _50___mcc_h5 = (_source1).cf49;
        let _51_cf49 = _50___mcc_h5;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else if (_source1.is_DC27) {
        let _52___mcc_h6 = (_source1).cf43;
        let _53_cf43 = _52___mcc_h6;
        return new _dafny.CodePoint('p'.codePointAt(0));
      } else {
        let _54___mcc_h7 = (_source1).cf50;
        let _55_cf50 = _54___mcc_h7;
        if (false) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }
      }
    };
    static fm35(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(568)), function (_56_i0) {
        return (_dafny.ZERO).minus(new BigNumber(-181));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-152)), function (_57_i1) {
        return new BigNumber(-182);
      })), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("eolvsmea"))).length)), _dafny.Seq.of(new BigNumber(673), new BigNumber(562))));
    };
    static fm36(globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(true))).Intersect((_dafny.MultiSet.fromElements(!(true))).Intersect(_dafny.MultiSet.fromElements(false)));
    };
    static fm37(p0, globalState) {
      return function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of _dafny.IntegerRange(new BigNumber(890), new BigNumber(345))) {
          let _58_v0 = _compr_12;
          if (((new BigNumber(890)).isLessThanOrEqualTo(_58_v0)) && ((_58_v0).isLessThan(new BigNumber(345)))) {
            _coll12.push([_module.__default.safeModuloInt(_58_v0, new BigNumber(137)),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("gmgrm")).length))]);
          }
        }
        return _coll12;
      }();
    };
    static fm38(p0, globalState) {
      let _source2 = _module.D5.create_DC15(_module.D5.create_DC12(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(-443))).length))));
      if (_source2.is_DC13) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_dafny.Map.Empty.slice().updateUnsafe((_module.D11.create_DC28(false)).dtor_cf44,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_module.D9.create_DC25(true, new BigNumber((_dafny.Set.fromElements(false)).length), false)).dtor_cf40)).length), new BigNumber(-792))).length))).length)))).Merge(function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).Elements) {
              let _59_v1 = _compr_14;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).contains(_59_v1)) {
                _coll14.push([_59_v1,new BigNumber(-606)]);
              }
            }
            return _coll14;
          }()).Keys.Elements) {
            let _60_v0 = _compr_13;
            if ((function () {
              let _coll15 = new _dafny.Map();
              for (const _compr_15 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).Elements) {
                let _59_v1 = _compr_15;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).contains(_59_v1)) {
                  _coll15.push([_59_v1,new BigNumber(-606)]);
                }
              }
              return _coll15;
            }()).contains(_60_v0)) {
              _coll13.push([_60_v0,_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(478),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(150)), function (_61_i0) {
                return new _dafny.CodePoint('t'.codePointAt(0));
              })).length))).length))]);
            }
          }
          return _coll13;
        }());
      } else if (_source2.is_DC14) {
        let _62___mcc_h0 = (_source2).cf23;
        let _63_cf23 = _62___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_63_cf23, _63_cf23, _63_cf23),_dafny.Map.Empty.slice().updateUnsafe(_63_cf23,new BigNumber(683)));
      } else if (_source2.is_DC12) {
        let _64___mcc_h1 = (_source2).cf22;
        let _65_cf22 = _64___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, false),_65_cf22);
      } else {
        let _66___mcc_h2 = (_source2).cf24;
        let _67_cf24 = _66___mcc_h2;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(164)))).Merge(function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false),_dafny.Seq.UnicodeFromString("uevfjb"))).Keys.Elements) {
            let _68_v2 = _compr_16;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false),_dafny.Seq.UnicodeFromString("uevfjb"))).contains(_68_v2)) {
              _coll16.push([_68_v2,_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber(-902)))]);
            }
          }
          return _coll16;
        }());
      }
    };
    static fm39(p0, globalState) {
      let _source3 = _module.D11.create_DC28(true);
      if (_source3.is_DC28) {
        let _69___mcc_h0 = (_source3).cf44;
        let _70_cf44 = _69___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(320),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), function (_71_i0) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), function (_72_i1) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length);
        }))).cardinality())))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(402),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("vavsi")).length))).cardinality())));
      } else if (_source3.is_DC29) {
        let _73___mcc_h1 = (_source3).cf45;
        let _74___mcc_h2 = (_source3).cf46;
        let _75___mcc_h3 = (_source3).cf47;
        let _76___mcc_h4 = (_source3).cf48;
        let _77_cf48 = _76___mcc_h4;
        let _78_cf47 = _75___mcc_h3;
        let _79_cf46 = _74___mcc_h2;
        let _80_cf45 = _73___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(_78_cf47,_78_cf47);
      } else if (_source3.is_DC30) {
        let _81___mcc_h5 = (_source3).cf49;
        let _82_cf49 = _81___mcc_h5;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(109),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_82_cf49,false)).length));
      } else if (_source3.is_DC27) {
        let _83___mcc_h6 = (_source3).cf43;
        let _84_cf43 = _83___mcc_h6;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(80),new BigNumber((_dafny.Seq.of(new BigNumber(408))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false, false)).length),new BigNumber(172)));
      } else {
        let _85___mcc_h7 = (_source3).cf50;
        let _86_cf50 = _85___mcc_h7;
        return function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_dafny.MultiSet.fromElements(new BigNumber(698))).Elements) {
            let _87_v0 = _compr_17;
            if ((_dafny.MultiSet.fromElements(new BigNumber(698))).contains(_87_v0)) {
              _coll17.push([_module.__default.safeDivisionInt(_87_v0, new BigNumber(-374)),new BigNumber(531)]);
            }
          }
          return _coll17;
        }();
      }
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("aynax")).length)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_88_i0) {
        return (_dafny.ZERO).minus(new BigNumber(-406));
      })));
    };
    static fm41(p0, globalState) {
      let _source4 = _module.D11.create_DC29(new BigNumber(-157), new BigNumber(677), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(48)), function (_89_i0) {
  return new BigNumber(-188);
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), function (_90_i1) {
  return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length));
}))).length), new BigNumber((_dafny.Seq.UnicodeFromString("tsx")).length))).length))).cardinality()), new BigNumber(681));
      if (_source4.is_DC28) {
        let _91___mcc_h0 = (_source4).cf44;
        let _92_cf44 = _91___mcc_h0;
        return _module.D4.create_DC11(_92_cf44);
      } else if (_source4.is_DC29) {
        let _93___mcc_h1 = (_source4).cf45;
        let _94___mcc_h2 = (_source4).cf46;
        let _95___mcc_h3 = (_source4).cf47;
        let _96___mcc_h4 = (_source4).cf48;
        let _97_cf48 = _96___mcc_h4;
        let _98_cf47 = _95___mcc_h3;
        let _99_cf46 = _94___mcc_h2;
        let _100_cf45 = _93___mcc_h1;
        return _module.D4.create_DC11(false);
      } else if (_source4.is_DC30) {
        let _101___mcc_h5 = (_source4).cf49;
        let _102_cf49 = _101___mcc_h5;
        return _module.D4.create_DC11(_102_cf49);
      } else if (_source4.is_DC27) {
        let _103___mcc_h6 = (_source4).cf43;
        let _104_cf43 = _103___mcc_h6;
        return _module.D4.create_DC11(false);
      } else {
        let _105___mcc_h7 = (_source4).cf50;
        let _106_cf50 = _105___mcc_h7;
        return _module.D4.create_DC11(false);
      }
    };
    static fm42(globalState) {
      return _module.D0.create_DC1((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fgdv")).length))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(-820), new BigNumber((_dafny.Seq.of(!(true))).length), new BigNumber(103), new BigNumber((_dafny.Seq.UnicodeFromString("aheg")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qehktwkga")).length))).length))).length))), (new BigNumber(-193)).plus(new BigNumber((_dafny.Seq.of(false, !(!(!(true))))).length)));
    };
    static fm43(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("yaynwj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), function (_107_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("yrbinxrmv"))).Difference((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.UnicodeFromString("fmj"))).Difference((_module.D12.create_DC32(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("mtvdgkj"), _dafny.Seq.UnicodeFromString("xxyl"))))).dtor_cf51));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), function (_108_i0) {
        return _module.D1.create_DC2((_module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), function (_109_i1) {
  return _dafny.MultiSet.fromElements(true);
})).length),new BigNumber(894)))).dtor_cf3);
      });
    };
    static fm45(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(false, false));
    };
    static fm46(p0, globalState) {
      return _dafny.Set.fromElements(function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(821), new BigNumber(144))) {
          let _110_v0 = _compr_18;
          if (((new BigNumber(821)).isLessThanOrEqualTo(_110_v0)) && ((_110_v0).isLessThan(new BigNumber(144)))) {
            _coll18.push([(_110_v0).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dav")).length))),new BigNumber(-163)]);
          }
        }
        return _coll18;
      }(), (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of _dafny.IntegerRange(new BigNumber(767), new BigNumber(440))) {
          let _111_v1 = _compr_19;
          if (((new BigNumber(767)).isLessThanOrEqualTo(_111_v1)) && ((_111_v1).isLessThan(new BigNumber(440)))) {
            _coll19.push([_module.__default.safeDivisionInt(_111_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D2.create_DC6(_dafny.Seq.of(new BigNumber(185), new BigNumber((_dafny.Seq.UnicodeFromString("kvwjkcv")).length), new BigNumber(685)))).dtor_cf11).length),new BigNumber((_dafny.Set.fromElements(new BigNumber(152), new BigNumber((_dafny.Set.fromElements(!(true))).length))).length))).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(70), new BigNumber(800))).cardinality())]);
          }
        }
        return _coll19;
      }()).length)),new BigNumber(489))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(955),new BigNumber(299))), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(383))).length),new BigNumber((_dafny.Seq.UnicodeFromString("yi")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_112_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length)),new BigNumber(120))));
    };
    static fm47(p0, p1, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), ((true) ? (new _dafny.CodePoint('i'.codePointAt(0))) : (new _dafny.CodePoint('l'.codePointAt(0)))), new _dafny.CodePoint('o'.codePointAt(0)));
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC6(_dafny.Seq.of(new BigNumber(794), new BigNumber(-102)));
    };
    static m0(p0, globalState) {
      let r0 = _dafny.ZERO;
      let _113_v0;
      _113_v0 = new BigNumber(-3);
      let _114_v1;
      _114_v1 = _dafny.Set.fromElements(_113_v0);
      let _115_v2;
      _115_v2 = false;
      let _116_v3;
      _116_v3 = _dafny.Seq.of(true, _115_v2);
      let _117_v4;
      let _nw0 = new _module.C2();
      _nw0.__ctor(_115_v2);
      _117_v4 = _nw0;
      let _118_v5;
      _118_v5 = _dafny.MultiSet.fromElements(_117_v4, _117_v4);
      let _119_v6;
      let _nw1 = new _module.C4();
      _nw1.__ctor(_115_v2, false);
      _119_v6 = _nw1;
      let _120_v7;
      _120_v7 = _dafny.Set.fromElements(_119_v6);
      let _121_v8;
      let _nw2 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _121_v8 = _nw2;
      let _122_v9;
      _122_v9 = _dafny.Map.Empty.slice().updateUnsafe(_121_v8,(_119_v6).f16);
      let _123_v10;
      _123_v10 = _122_v9;
      let _124_v11;
      _124_v11 = _dafny.Map.Empty.slice().updateUnsafe((_119_v6).f16,_115_v2);
      let _125_v12;
      _125_v12 = _module.D4.create_DC11((_119_v6).f16);
      let _126_v13;
      _126_v13 = _dafny.Map.Empty.slice().updateUnsafe(_113_v0,_113_v0);
      let _127_v14;
      _127_v14 = _dafny.Seq.UnicodeFromString("wm");
      let _128_v15;
      _128_v15 = new _dafny.CodePoint('v'.codePointAt(0));
      let _129_v16;
      _129_v16 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("yyl"), _dafny.Seq.update(_127_v14, _module.__default.safeIndex(_113_v0, new BigNumber((_127_v14).length)), _128_v15));
      let _130_v17;
      _130_v17 = _dafny.MultiSet.fromElements(_116_v3);
      let _131_v18;
      let _nw3 = Array((new BigNumber(29)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = (_113_v0).isLessThan(new BigNumber((_114_v1).length));
      _nw3[(_dafny.ONE).toNumber()] = (_116_v3)[_module.__default.safeIndex(new BigNumber((_118_v5).cardinality()), new BigNumber((_116_v3).length))];
      _nw3[(new BigNumber(2)).toNumber()] = (_120_v7).IsProperSubsetOf(_120_v7);
      _nw3[(new BigNumber(3)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(4)).toNumber()] = !((_119_v6).f16);
      _nw3[(new BigNumber(5)).toNumber()] = (_119_v6).f16;
      _nw3[(new BigNumber(6)).toNumber()] = (_119_v6).f16;
      _nw3[(new BigNumber(7)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(8)).toNumber()] = true;
      _nw3[(new BigNumber(9)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(10)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(11)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(12)).toNumber()] = (_119_v6).f16;
      _nw3[(new BigNumber(13)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_121_v8,_115_v2))).equals((_122_v9).update(_121_v8, (_119_v6).f16));
      _nw3[(new BigNumber(14)).toNumber()] = (((_122_v9).contains(_121_v8)) ? ((_122_v9).get(_121_v8)) : (_115_v2));
      _nw3[(new BigNumber(15)).toNumber()] = (_115_v2) === (_115_v2);
      _nw3[(new BigNumber(16)).toNumber()] = (_module.__default.fm25(_114_v1, _124_v11, _113_v0, _115_v2, globalState)).IsDisjointFrom(_module.__default.fm25(_114_v1, _module.__default.fm20(_115_v2, _113_v0, (_119_v6).f16, globalState), _113_v0, (_119_v6).f16, globalState));
      _nw3[(new BigNumber(17)).toNumber()] = _115_v2;
      _nw3[(new BigNumber(18)).toNumber()] = (_119_v6).f16;
      _nw3[(new BigNumber(19)).toNumber()] = (_119_v6).f16;
      _nw3[(new BigNumber(20)).toNumber()] = false;
      _nw3[(new BigNumber(21)).toNumber()] = (_113_v0).isLessThanOrEqualTo(new BigNumber(-421));
      _nw3[(new BigNumber(22)).toNumber()] = (((_119_v6).f16) ? (_115_v2) : ((_119_v6).f16));
      _nw3[(new BigNumber(23)).toNumber()] = !((_125_v12).dtor_cf21);
      _nw3[(new BigNumber(24)).toNumber()] = _module.__default.fm0(_126_v13, _113_v0, globalState);
      _nw3[(new BigNumber(25)).toNumber()] = (_129_v16).IsSubsetOf(_dafny.MultiSet.fromElements(_127_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), function (_132_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _127_v14));
      _nw3[(new BigNumber(26)).toNumber()] = (!(false)) || (_115_v2);
      _nw3[(new BigNumber(27)).toNumber()] = !((((_119_v6).f16) ? (_115_v2) : ((_119_v6).f16)));
      _nw3[(new BigNumber(28)).toNumber()] = (_113_v0).isLessThanOrEqualTo((((_130_v17).contains(_116_v3)) ? ((_130_v17).get(_116_v3)) : (_113_v0)));
      _131_v18 = _nw3;
      _131_v18 = _131_v18;
      let _133_v19;
      _133_v19 = _dafny.Seq.of(_116_v3);
      let _source5 = (_133_v19)[_module.__default.safeIndex(_113_v0, new BigNumber((_133_v19).length))];
      let _134___mcc_h0 = _source5;
      let _135_cf34 = _134___mcc_h0;
      r0 = _113_v0;
      _115_v2 = (_119_v6).f16;
      let _136_v20;
      let _nw4 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
      _136_v20 = _nw4;
      let _137_v21;
      _137_v21 = _dafny.MultiSet.fromElements(_128_v15);
      let _index0 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_136_v20).length));
      (_136_v20)[_index0] = (_137_v21).Intersect(_137_v21);
      let _index1 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_136_v20).length));
      let _rhs0 = (_137_v21).Union(_dafny.MultiSet.fromElements(_128_v15, new _dafny.CodePoint('s'.codePointAt(0))));
      let _rhs1 = !(_115_v2);
      let _lhs0 = _136_v20;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_136_v20).length));
      _lhs0[_lhs1] = _rhs0;
      _115_v2 = _rhs1;
      _115_v2 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rokcgkova")).length))).isLessThanOrEqualTo(_113_v0);
      let _138_v22;
      _138_v22 = _module.D1.create_DC3(new BigNumber(104));
      let _139_i1;
      _139_i1 = _dafny.ZERO;
      L0: {
        let _pat_let_tv0 = _119_v6;
        let _pat_let_tv1 = _119_v6;
        let _pat_let_tv2 = _119_v6;
        let _pat_let_tv3 = _115_v2;
        let _pat_let_tv4 = _119_v6;
        let _pat_let_tv5 = _113_v0;
        while (function (_source6) {
          if (_source6.is_DC3) {
            let _144___mcc_h1 = (_source6).cf4;
            let _145_cf4 = _144___mcc_h1;
            return true;
          } else if (_source6.is_DC4) {
            let _146___mcc_h2 = (_source6).cf5;
            let _147___mcc_h3 = (_source6).cf6;
            let _148___mcc_h4 = (_source6).cf7;
            let _149_cf7 = _148___mcc_h4;
            let _150_cf6 = _147___mcc_h3;
            let _151_cf5 = _146___mcc_h2;
            return (_dafny.Set.fromElements((_pat_let_tv0).f16)).IsSubsetOf(_dafny.Set.fromElements((_pat_let_tv1).f16));
          } else if (_source6.is_DC5) {
            let _152___mcc_h5 = (_source6).cf8;
            let _153___mcc_h6 = (_source6).cf9;
            let _154___mcc_h7 = (_source6).cf10;
            let _155_cf10 = _154___mcc_h7;
            let _156_cf9 = _153___mcc_h6;
            let _157_cf8 = _152___mcc_h5;
            return (_pat_let_tv2).f16;
          } else {
            let _158___mcc_h8 = (_source6).cf3;
            let _159_cf3 = _158___mcc_h8;
            return (new BigNumber((_dafny.Set.fromElements(_pat_let_tv3, (_pat_let_tv4).f16)).length)).isEqualTo(_pat_let_tv5);
          }
        }(((_115_v2) ? (_138_v22) : (_138_v22)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_139_i1)) {
              break L0;
            }
            _139_i1 = (_139_i1).plus(_dafny.ONE);
            let _index2 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_121_v8).length));
            (_121_v8)[_index2] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nntti"), _127_v14), _module.__default.safeIndex(_113_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nntti"), _127_v14)).length)), _128_v15)).length);
            let _index3 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_121_v8).length));
            (_121_v8)[_index3] = (_113_v0).plus(_113_v0);
            let _140_v23;
            let _nw5 = new _module.C4();
            _nw5.__ctor(!((_121_v8)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_121_v8).length))]).isEqualTo((_dafny.ZERO).minus((_121_v8)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_121_v8).length))])), (_114_v1).IsSubsetOf((_module.D11.create_DC27(_114_v1)).dtor_cf43));
            _140_v23 = _nw5;
            let _141_v24;
            _141_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_113_v0),_131_v18);
            let _142_v25;
            _142_v25 = _module.D1.create_DC4(_131_v18, new BigNumber((_127_v14).length), _128_v15);
            _131_v18 = (((_141_v24).contains(_113_v0)) ? ((_141_v24).get(_113_v0)) : ((_142_v25).dtor_cf5));
            let _143_v26;
            let _nw6 = Array((new BigNumber(13)).toNumber()).fill([]);
            _143_v26 = _nw6;
            let _index4 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_143_v26).length));
            (_143_v26)[_index4] = _131_v18;
            let _index5 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_143_v26).length));
            (_143_v26)[_index5] = _131_v18;
          }
        }
      }
      let _160_v27;
      _160_v27 = _dafny.Set.fromElements(_124_v11, _124_v11, _124_v11);
      _115_v2 = (_160_v27).equals(_160_v27);
      let _161_v28;
      let _nw7 = Array((new BigNumber(22)).toNumber()).fill(_module.D7.Default());
      _161_v28 = _nw7;
      let _162_v29;
      _162_v29 = _module.D7.create_DC20(new BigNumber(-722), _128_v15);
      let _index6 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_161_v28).length));
      (_161_v28)[_index6] = _162_v29;
      let _index7 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_161_v28).length));
      (_161_v28)[_index7] = _162_v29;
      _113_v0 = _113_v0;
      r0 = _113_v0;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _163_v0;
      _163_v0 = false;
      let _164_v1;
      _164_v1 = _dafny.Seq.of(_163_v0, !(_163_v0), _163_v0, !(_163_v0));
      let _165_v2;
      _165_v2 = new BigNumber(-199);
      let _166_v3;
      let _nw8 = Array((new BigNumber(27)).toNumber());
      _nw8[(_dafny.ZERO).toNumber()] = _163_v0;
      _nw8[(_dafny.ONE).toNumber()] = _163_v0;
      _nw8[(new BigNumber(2)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(3)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(4)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(5)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(6)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(7)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(8)).toNumber()] = true;
      _nw8[(new BigNumber(9)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(10)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(11)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(12)).toNumber()] = !(_163_v0);
      _nw8[(new BigNumber(13)).toNumber()] = true;
      _nw8[(new BigNumber(14)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(15)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(16)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(17)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(18)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(19)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(20)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(21)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(22)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(23)).toNumber()] = (_164_v1)[_module.__default.safeIndex(_165_v2, new BigNumber((_164_v1).length))];
      _nw8[(new BigNumber(24)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(25)).toNumber()] = _163_v0;
      _nw8[(new BigNumber(26)).toNumber()] = _163_v0;
      _166_v3 = _nw8;
      let _167_v4;
      _167_v4 = _module.D0.create_DC0(false);
      let _168_v5;
      _168_v5 = _dafny.Set.fromElements(_164_v1, _dafny.Seq.update(_dafny.Seq.of(_163_v0, (_167_v4).dtor_cf0, !(false), _163_v0, _163_v0), _module.__default.safeIndex(_165_v2, new BigNumber((_dafny.Seq.of(_163_v0, (_167_v4).dtor_cf0, !(false), _163_v0, _163_v0)).length)), false), _164_v1);
      let _169_v6;
      _169_v6 = _dafny.Seq.UnicodeFromString("nftgk");
      let _170_globalState;
      let _nw9 = new _module.GlobalState();
      _nw9.__ctor(new BigNumber(-586), _166_v3, true, false, _164_v1, (_168_v5).Union(_168_v5), false, _dafny.MultiSet.fromElements(_165_v2), true, new BigNumber(721), _166_v3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), function (_171_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _169_v6), new BigNumber(604));
      _170_globalState = _nw9;
      let _172_v7;
      _172_v7 = _module.D0.create_DC1(_163_v0, _module.__default.safeDivisionInt(_165_v2, _165_v2));
      let _index8 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
      (_166_v3)[_index8] = _dafny.Seq.IsProperPrefixOf(_169_v6, _169_v6);
      let _173_v8;
      _173_v8 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_165_v2),new BigNumber(833));
      let _174_v9;
      _174_v9 = _dafny.MultiSet.fromElements(_165_v2);
      let _175_v10;
      _175_v10 = new _dafny.CodePoint('d'.codePointAt(0));
      let _index9 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
      let _rhs2 = _172_v7;
      let _rhs3 = _module.__default.fm0(_173_v8, (_dafny.ZERO).minus(new BigNumber((_174_v9).cardinality())), _170_globalState);
      let _rhs4 = _dafny.Seq.contains(_dafny.Seq.Concat(_169_v6, _dafny.Seq.UnicodeFromString("qtgkj")), _175_v10);
      let _rhs5 = _163_v0;
      let _rhs6 = _166_v3;
      let _lhs2 = _166_v3;
      let _lhs3 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
      _172_v7 = _rhs2;
      _lhs2[_lhs3] = _rhs3;
      _163_v0 = _rhs4;
      _163_v0 = _rhs5;
      _166_v3 = _rhs6;
      let _176_i1;
      _176_i1 = _dafny.ZERO;
      L1: {
        while (false) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_176_i1)) {
              break L1;
            }
            _176_i1 = (_176_i1).plus(_dafny.ONE);
            let _177_v11;
            let _init0 = ((_178_v1) => function (_179_i2) {
              return _178_v1;
            })(_164_v1);
            let _nw10 = Array((new BigNumber(13)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw10.length); _i0_0++) {
              _nw10[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _177_v11 = _nw10;
            let _index10 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_177_v11).length));
            (_177_v11)[_index10] = _164_v1;
            let _index11 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_177_v11).length));
            (_177_v11)[_index11] = _164_v1;
            let _180_v12;
            let _nw11 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
            _180_v12 = _nw11;
            let _181_v13;
            _181_v13 = _dafny.Map.Empty.slice().updateUnsafe(_166_v3,_module.__default.fm0(_173_v8, _165_v2, _170_globalState));
            let _182_v14;
            _182_v14 = _dafny.Map.Empty.slice().updateUnsafe(_163_v0,(((_181_v13).contains(_166_v3)) ? ((_181_v13).get(_166_v3)) : (true)));
            let _183_v15;
            let _nw12 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _183_v15 = _nw12;
            let _184_v16;
            _184_v16 = _dafny.Set.fromElements(_183_v15, _183_v15);
            let _185_v17;
            _185_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_182_v14).length),_184_v16);
            let _index12 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_180_v12).length));
            (_180_v12)[_index12] = ((((_185_v17).contains(new BigNumber(-75))) ? ((_185_v17).get(new BigNumber(-75))) : (_184_v16))).Union(_184_v16);
            let _index13 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_180_v12).length));
            (_180_v12)[_index13] = _dafny.Set.fromElements(_183_v15);
            _175_v10 = _175_v10;
            let _186_v18;
            _186_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wm"),(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
            let _187_v19;
            let _out0;
            _out0 = _module.__default.m0(_186_v18, _170_globalState);
            _187_v19 = _out0;
          }
        }
      }
      let _188_v20;
      let _init1 = ((_189_v2) => function (_190_i3) {
        return (_190_i3).plus(_189_v2);
      })(_165_v2);
      let _nw13 = Array((new BigNumber(27)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw13.length); _i0_1++) {
        _nw13[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _188_v20 = _nw13;
      let _191_v21;
      _191_v21 = _dafny.Map.Empty.slice().updateUnsafe(_165_v2,_188_v20);
      _191_v21 = _191_v21;
      let _192_v22;
      _192_v22 = _module.D1.create_DC2(_173_v8);
      let _193_i4;
      _193_i4 = _dafny.ZERO;
      L2: {
        while (_module.__default.fm0((_192_v22).dtor_cf3, _165_v2, _170_globalState)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_193_i4)) {
              break L2;
            }
            _193_i4 = (_193_i4).plus(_dafny.ONE);
            let _index14 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
            (_166_v3)[_index14] = _module.__default.fm0(_173_v8, _165_v2, _170_globalState);
            let _194_v23;
            _194_v23 = _dafny.Map.Empty.slice().updateUnsafe(_165_v2,!(_163_v0));
            let _195_v24;
            _195_v24 = _dafny.Seq.of(_165_v2);
            let _196_v25;
            _196_v25 = _dafny.Map.Empty.slice().updateUnsafe(_194_v23,_195_v24);
            let _197_v26;
            _197_v26 = _dafny.Set.fromElements((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))], _163_v0, _163_v0, !(false));
            let _198_v27;
            _198_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_197_v26).length),_195_v24);
            let _199_v28;
            _199_v28 = _dafny.Map.Empty.slice().updateUnsafe((_196_v25).Merge(_dafny.Map.Empty.slice().updateUnsafe(_194_v23,(((_198_v27).contains(new BigNumber((_169_v6).length))) ? ((_198_v27).get(new BigNumber((_169_v6).length))) : (_195_v24)))),_module.__default.fm0(_173_v8, _165_v2, _170_globalState));
            _199_v28 = (_199_v28).update((_196_v25).Merge(_dafny.Map.Empty.slice().updateUnsafe(_194_v23,_195_v24)), _163_v0);
            let _index15 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_188_v20).length));
            (_188_v20)[_index15] = new BigNumber((function () {
              let _coll20 = new _dafny.Set();
              for (const _compr_20 of _dafny.IntegerRange(new BigNumber(448), new BigNumber(86))) {
                let _200_v29 = _compr_20;
                if (((new BigNumber(448)).isLessThanOrEqualTo(_200_v29)) && ((_200_v29).isLessThan(new BigNumber(86)))) {
                  _coll20.add((_200_v29).multipliedBy(_165_v2));
                }
              }
              return _coll20;
            }()).length);
            let _201_v30;
            _201_v30 = _dafny.Map.Empty.slice().updateUnsafe(_163_v0,new BigNumber((_195_v24).length));
            let _index16 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_188_v20).length));
            let _index17 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
            let _rhs7 = false;
            let _rhs8 = _165_v2;
            let _rhs9 = _188_v20;
            let _rhs10 = (new BigNumber(782)).plus(_165_v2);
            let _rhs11 = !((_201_v30).Merge(_201_v30)).contains((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
            let _lhs4 = _188_v20;
            let _lhs5 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_188_v20).length));
            let _lhs6 = _166_v3;
            let _lhs7 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
            _163_v0 = _rhs7;
            _lhs4[_lhs5] = _rhs8;
            _188_v20 = _rhs9;
            _165_v2 = _rhs10;
            _lhs6[_lhs7] = _rhs11;
            let _202_v31;
            _202_v31 = _dafny.Seq.of(_166_v3, _166_v3);
            let _index18 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
            (_166_v3)[_index18] = !_dafny.Seq.contains(_202_v31, _166_v3);
          }
        }
      }
      let _203_v32;
      _203_v32 = _dafny.Map.Empty.slice().updateUnsafe(!((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]),_165_v2);
      let _hi0 = _165_v2;
      for (let _204_i5 = new BigNumber(((_203_v32).Merge(_203_v32)).length); _204_i5.isLessThan(_hi0); _204_i5 = _204_i5.plus(_dafny.ONE)) {
        let _205_v33;
        let _nw14 = Array((new BigNumber(21)).toNumber()).fill([]);
        _205_v33 = _nw14;
        let _206_v34;
        _206_v34 = _dafny.Seq.of(_165_v2);
        let _207_v35;
        _207_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_206_v34, _206_v34),((_163_v0) ? (true) : (_163_v0)));
        let _208_v36;
        _208_v36 = _dafny.Map.Empty.slice().updateUnsafe(_188_v20,_165_v2);
        let _index19 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        let _index20 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        let _rhs12 = _165_v2;
        let _rhs13 = _205_v33;
        let _rhs14 = _module.__default.fm1(false, _172_v7, _170_globalState);
        let _rhs15 = (((((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]) ? ((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]) : (_163_v0))) ? (!(_208_v36).equals(_208_v36)) : (true));
        let _rhs16 = (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))];
        let _lhs8 = _166_v3;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        let _lhs10 = _166_v3;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        _165_v2 = _rhs12;
        _205_v33 = _rhs13;
        _207_v35 = _rhs14;
        _lhs8[_lhs9] = _rhs15;
        _lhs10[_lhs11] = _rhs16;
        _165_v2 = (_dafny.ZERO).minus((_204_i5).minus(_204_i5));
        let _index21 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index21] = !(_165_v2).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("mjidffht")).length));
        _163_v0 = ((_dafny.ZERO).minus(_204_i5)).isEqualTo(new BigNumber((_dafny.Seq.Concat(_169_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-249)), ((_209_v10) => function (_210_i6) {
          return _209_v10;
        })(_175_v10)))).length));
      }
      let _211_v37;
      _211_v37 = _dafny.Set.fromElements(_175_v10);
      let _212_v38;
      _212_v38 = _dafny.MultiSet.fromElements(_211_v37);
      let _213_v39;
      _213_v39 = _dafny.MultiSet.fromElements(false, (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))], (_164_v1)[_module.__default.safeIndex(new BigNumber((_174_v9).cardinality()), new BigNumber((_164_v1).length))]);
      let _214_v41;
      _214_v41 = _dafny.Seq.of(_211_v37, function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])).Keys.Elements) {
          let _215_v40 = _compr_21;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])).contains(_215_v40)) {
            _coll21.add(_215_v40);
          }
        }
        return _coll21;
      }());
      let _216_v42;
      let _nw15 = Array((new BigNumber(9)).toNumber());
      _nw15[(_dafny.ZERO).toNumber()] = _212_v38;
      _nw15[(_dafny.ONE).toNumber()] = _212_v38;
      _nw15[(new BigNumber(2)).toNumber()] = _module.__default.fm2(!(false), _165_v2, _170_globalState);
      _nw15[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm3(new BigNumber((_213_v39).cardinality()), _165_v2, _dafny.Seq.of(_174_v9, _174_v9), _175_v10, _170_globalState));
      _nw15[(new BigNumber(4)).toNumber()] = _212_v38;
      _nw15[(new BigNumber(5)).toNumber()] = _212_v38;
      _nw15[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_214_v41);
      _nw15[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(_211_v37, _211_v37)).update(_211_v37, _module.__default.abs(_165_v2));
      _nw15[(new BigNumber(8)).toNumber()] = _212_v38;
      _216_v42 = _nw15;
      let _index22 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_216_v42).length));
      (_216_v42)[_index22] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(389)), ((_217_v10) => function (_218_i7) {
        return function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe(_217_v10,new BigNumber(613))).Keys.Elements) {
            let _219_v43 = _compr_22;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_217_v10,new BigNumber(613))).contains(_219_v43)) {
              _coll22.add(_219_v43);
            }
          }
          return _coll22;
        }();
      })(_175_v10)));
      let _index23 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_216_v42).length));
      (_216_v42)[_index23] = (_212_v38).update(_211_v37, _module.__default.abs(_165_v2));
      let _hi1 = (_165_v2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_165_v2,_165_v2)).length));
      for (let _220_i8 = _165_v2; _220_i8.isLessThan(_hi1); _220_i8 = _220_i8.plus(_dafny.ONE)) {
        _188_v20 = _188_v20;
        let _index24 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index24] = !(!((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]));
        _165_v2 = _220_i8;
        if ((_165_v2).isLessThanOrEqualTo(new BigNumber(452))) {
          _165_v2 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_164_v1, _module.__default.safeIndex(new BigNumber((_169_v6).length), new BigNumber((_164_v1).length)), (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]), _dafny.Seq.of(!(_163_v0), _163_v0, _163_v0, _163_v0))).length)).plus(_220_i8);
          let _221_v44;
          _221_v44 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_165_v2),(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
          let _222_v46;
          let _nw16 = new _module.C2();
          _nw16.__ctor(!(_221_v44).equals(function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(603), new BigNumber(37))) {
              let _223_v45 = _compr_23;
              if (((new BigNumber(603)).isLessThanOrEqualTo(_223_v45)) && ((_223_v45).isLessThan(new BigNumber(37)))) {
                _coll23.push([(_223_v45).plus(_165_v2),_163_v0]);
              }
            }
            return _coll23;
          }()));
          _222_v46 = _nw16;
          let _index25 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          (_166_v3)[_index25] = !(_163_v0);
          _163_v0 = (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))];
          let _224_v47;
          let _nw17 = Array((new BigNumber(16)).toNumber()).fill([]);
          _224_v47 = _nw17;
          let _225_v48;
          let _nw18 = Array((new BigNumber(11)).toNumber());
          _225_v48 = _nw18;
          let _index26 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_224_v47).length));
          (_224_v47)[_index26] = _225_v48;
          let _index27 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_224_v47).length));
          let _rhs17 = ((_165_v2).minus(new BigNumber(957))).isLessThanOrEqualTo(new BigNumber(855));
          let _rhs18 = _220_i8;
          let _rhs19 = (((_173_v8).contains(new BigNumber(((_221_v44).Merge(_221_v44)).length))) ? ((_173_v8).get(new BigNumber(((_221_v44).Merge(_221_v44)).length))) : (_220_i8));
          let _rhs20 = _225_v48;
          let _lhs12 = _224_v47;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_224_v47).length));
          _163_v0 = _rhs17;
          _165_v2 = _rhs18;
          _165_v2 = _rhs19;
          _lhs12[_lhs13] = _rhs20;
        } else {
          _169_v6 = _dafny.Seq.UnicodeFromString("tkuqi");
          let _226_v49;
          _226_v49 = _module.D2.create_DC8(_163_v0, _163_v0, _163_v0);
          let _index28 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          let _rhs21 = _163_v0;
          let _rhs22 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_220_i8).multipliedBy(new BigNumber(479)), _165_v2)));
          let _rhs23 = (_226_v49).dtor_cf17;
          let _lhs14 = _166_v3;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          _lhs14[_lhs15] = _rhs21;
          _165_v2 = _rhs22;
          _163_v0 = _rhs23;
          _175_v10 = new _dafny.CodePoint('l'.codePointAt(0));
          let _227_v50;
          let _nw19 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _227_v50 = _nw19;
          let _228_v51;
          _228_v51 = _dafny.Set.fromElements(_165_v2, _165_v2);
          let _index29 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_227_v50).length));
          (_227_v50)[_index29] = _228_v51;
          let _index30 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_227_v50).length));
          (_227_v50)[_index30] = _228_v51;
          let _229_v52;
          let _nw20 = new _module.C5();
          _nw20.__ctor(!(_163_v0), _188_v20);
          _229_v52 = _nw20;
          _229_v52 = _229_v52;
        }
      }
      let _230_i9;
      _230_i9 = _dafny.ZERO;
      L3: {
        while (!((((_213_v39).contains(_163_v0)) ? ((_213_v39).get(_163_v0)) : (new BigNumber(-671)))).isEqualTo(_165_v2)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_230_i9)) {
              break L3;
            }
            _230_i9 = (_230_i9).plus(_dafny.ONE);
            let _index31 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_188_v20).length));
            (_188_v20)[_index31] = _165_v2;
            let _index32 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_188_v20).length));
            (_188_v20)[_index32] = (((_203_v32).contains((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])) ? ((_203_v32).get((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])) : (_165_v2));
            let _231_v53;
            _231_v53 = _dafny.Map.Empty.slice().updateUnsafe(_169_v6,(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
            let _232_v54;
            let _out1;
            _out1 = _module.__default.m0(_231_v53, _170_globalState);
            _232_v54 = _out1;
            let _index33 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
            (_166_v3)[_index33] = _163_v0;
            let _index34 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_188_v20).length));
            let _rhs24 = new BigNumber(711);
            let _rhs25 = (_232_v54).plus(new BigNumber(-301));
            let _lhs16 = _188_v20;
            let _lhs17 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_188_v20).length));
            _165_v2 = _rhs24;
            _lhs16[_lhs17] = _rhs25;
          }
        }
      }
      let _233_v55;
      let _nw21 = new _module.C3();
      _nw21.__ctor(_163_v0, !((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]));
      _233_v55 = _nw21;
      _233_v55 = ((!_dafny.Seq.contains(_164_v1, (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])) ? (_233_v55) : (_233_v55));
      let _234_v56;
      _234_v56 = _dafny.Seq.of(_188_v20, _188_v20, (((_233_v55).f19) ? (_188_v20) : (_188_v20)));
      _234_v56 = _234_v56;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_188_v20).length))) {
        let _235_i10 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_235_i10)) && ((_235_i10).isLessThan(new BigNumber((_188_v20).length))))) {
          (_188_v20)[(_235_i10)] = _module.__default.safeDivisionInt(_235_i10, _165_v2);
        }
      }
      if (_163_v0) {
        let _index35 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index35] = true;
        let _236_v57;
        let _out2;
        _out2 = (_233_v55).m1(_module.__default.fm31(_170_globalState), _170_globalState);
        _236_v57 = _out2;
        let _237_v58;
        let _init2 = ((_238_v55) => function (_239_i11) {
          return _dafny.Set.fromElements((_238_v55).f19);
        })(_233_v55);
        let _nw22 = Array((new BigNumber(29)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
          _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _237_v58 = _nw22;
        let _240_v60;
        _240_v60 = _dafny.Set.fromElements((_233_v55).f19, (_233_v55).f19, (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))], !(_module.__default.fm0(function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (_174_v9).Elements) {
            let _241_v59 = _compr_24;
            if ((_174_v9).contains(_241_v59)) {
              _coll24.push([(_241_v59).plus(_165_v2),_165_v2]);
            }
          }
          return _coll24;
        }(), _165_v2, _170_globalState)));
        let _index36 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_237_v58).length));
        (_237_v58)[_index36] = (((_164_v1)[_module.__default.safeIndex(_165_v2, new BigNumber((_164_v1).length))]) ? (_240_v60) : (_dafny.Set.fromElements((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])));
        let _index37 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_237_v58).length));
        (_237_v58)[_index37] = (_240_v60).Union(_240_v60);
        let _242_v61;
        let _nw23 = new _module.C3();
        _nw23.__ctor((_233_v55).f19, (_165_v2).isLessThan(new BigNumber((_dafny.Set.fromElements(_165_v2)).length)));
        _242_v61 = _nw23;
        let _index38 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index38] = (_module.__default.safeModuloInt(_165_v2, _165_v2)).isLessThan(new BigNumber(((_240_v60).Difference((_237_v58)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_237_v58).length))])).length));
      } else {
        let _243_v63;
        _243_v63 = _dafny.Seq.of(_165_v2, _165_v2);
        let _pat_let_tv6 = _165_v2;
        let _244_v64;
        _244_v64 = _dafny.Map.Empty.slice().updateUnsafe((_233_v55).fm4(_163_v0, function (_pat_let0_0) {
          return function (_245_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_246_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_245_dt__update__tmp_h0).dtor_cf1, _246_dt__update_hcf2_h0);
              }(_pat_let1_0);
            }(_pat_let_tv6);
          }(_pat_let0_0);
        }(_172_v7), new BigNumber((_203_v32).length), function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_243_v63).Elements) {
            let _247_v62 = _compr_25;
            if (_dafny.Seq.contains(_243_v63, _247_v62)) {
              _coll25.push([(_247_v62).minus(_165_v2),_165_v2]);
            }
          }
          return _coll25;
        }(), _170_globalState),_203_v32);
        _165_v2 = (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_233_v55).f19,_module.__default.fm31(_170_globalState))).Merge((((_244_v64).contains((_233_v55).f19)) ? ((_244_v64).get((_233_v55).f19)) : (_203_v32)))).length));
        let _248_v65;
        let _nw24 = Array((new BigNumber(4)).toNumber());
        _248_v65 = _nw24;
        _248_v65 = _248_v65;
        let _index39 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index39] = (_233_v55).f19;
        _165_v2 = _165_v2;
        _174_v9 = _174_v9;
      }
      let _249_v66;
      let _nw25 = Array((new BigNumber(14)).toNumber()).fill([]);
      _249_v66 = _nw25;
      let _index40 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
      (_249_v66)[_index40] = _188_v20;
      let _index41 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
      (_249_v66)[_index41] = _188_v20;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_166_v3).length))) {
        let _250_i12 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_250_i12)) && ((_250_i12).isLessThan(new BigNumber((_166_v3).length))))) {
          (_166_v3)[(_250_i12)] = !(_163_v0) || ((((_233_v55).f19) ? ((_233_v55).f19) : ((_233_v55).f19)));
        }
      }
      let _251_v67;
      _251_v67 = _module.D5.create_DC13();
      let _source7 = _251_v67;
      if (_source7.is_DC13) {
        let _252_v68;
        _252_v68 = _dafny.Set.fromElements(_165_v2);
        let _index42 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        let _rhs26 = (_165_v2).minus(_165_v2);
        let _rhs27 = (_252_v68).equals(_dafny.Set.fromElements(_165_v2, new BigNumber((_169_v6).length)));
        let _lhs18 = _166_v3;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        _165_v2 = _rhs26;
        _lhs18[_lhs19] = _rhs27;
        let _253_v69;
        let _nw26 = Array((new BigNumber(8)).toNumber()).fill(_module.D5.Default());
        _253_v69 = _nw26;
        let _254_v70;
        _254_v70 = _dafny.Map.Empty.slice().updateUnsafe(_233_v55,_173_v8);
        let _255_v71;
        _255_v71 = _dafny.Seq.of(_173_v8);
        let _pat_let_tv7 = _165_v2;
        let _pat_let_tv8 = _165_v2;
        let _256_v72;
        let _nw27 = Array((new BigNumber(23)).toNumber());
        _nw27[(_dafny.ZERO).toNumber()] = (((_254_v70).contains(_233_v55)) ? ((_254_v70).get(_233_v55)) : (_173_v8));
        _nw27[(_dafny.ONE).toNumber()] = _173_v8;
        _nw27[(new BigNumber(2)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(3)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(4)).toNumber()] = (_173_v8).update(_165_v2, _165_v2);
        _nw27[(new BigNumber(5)).toNumber()] = (((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]) ? ((_255_v71)[_module.__default.safeIndex(_165_v2, new BigNumber((_255_v71).length))]) : (_173_v8));
        _nw27[(new BigNumber(6)).toNumber()] = ((_163_v0) ? (_173_v8) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(303),new BigNumber(904))));
        _nw27[(new BigNumber(7)).toNumber()] = ((function (_pat_let2_0) {
          return function (_257_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_258_dt__update_hcf3_h0) {
                return _module.D1.create_DC2(_258_dt__update_hcf3_h0);
              }(_pat_let3_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv7,_pat_let_tv8));
          }(_pat_let2_0);
        }(_module.D1.create_DC2(_173_v8))).dtor_cf3).Merge(_173_v8);
        _nw27[(new BigNumber(8)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(9)).toNumber()] = _module.__default.fm39(_165_v2, _170_globalState);
        _nw27[(new BigNumber(10)).toNumber()] = ((_173_v8).update(new BigNumber((_module.__default.fm28(new BigNumber(203), _165_v2, !(true), _170_globalState)).length), _165_v2)).Merge(_173_v8);
        _nw27[(new BigNumber(11)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(12)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(13)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(14)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(15)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(16)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(17)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(18)).toNumber()] = ((_163_v0) ? (_173_v8) : (_173_v8));
        _nw27[(new BigNumber(19)).toNumber()] = _module.__default.fm39(_165_v2, _170_globalState);
        _nw27[(new BigNumber(20)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(21)).toNumber()] = _173_v8;
        _nw27[(new BigNumber(22)).toNumber()] = _173_v8;
        _256_v72 = _nw27;
        let _index43 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_256_v72).length));
        (_256_v72)[_index43] = _173_v8;
        let _259_v73;
        _259_v73 = _dafny.Seq.of(_165_v2, _165_v2);
        let _260_v74;
        _260_v74 = _dafny.Set.fromElements(_259_v73, _259_v73, _259_v73, _259_v73);
        let _index44 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length));
        (_188_v20)[_index44] = new BigNumber((_260_v74).length);
        let _261_v75;
        _261_v75 = _dafny.Map.Empty.slice().updateUnsafe((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))],(_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
        let _262_v76;
        _262_v76 = _module.D4.create_DC11((_233_v55).f19);
        let _263_v77;
        _263_v77 = _dafny.Set.fromElements((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))], !((_262_v76).dtor_cf21));
        let _index45 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_256_v72).length));
        let _index46 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length));
        let _rhs28 = ((!(_261_v75).contains((_164_v1)[_module.__default.safeIndex(_165_v2, new BigNumber((_164_v1).length))])) ? (_253_v69) : (_253_v69));
        let _rhs29 = (_173_v8).update(new BigNumber((_174_v9).cardinality()), new BigNumber(((((_233_v55).f19) ? (_259_v73) : (_259_v73))).length));
        let _rhs30 = (new BigNumber(-208)).minus(new BigNumber(((_263_v77).Union(_dafny.Set.fromElements((_233_v55).f19))).length));
        let _lhs20 = _256_v72;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_256_v72).length));
        let _lhs22 = _188_v20;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length));
        _253_v69 = _rhs28;
        _lhs20[_lhs21] = _rhs29;
        _lhs22[_lhs23] = _rhs30;
        if ((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]) {
          let _264_v78;
          _264_v78 = _dafny.Seq.of(_259_v73, (((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]) ? (_259_v73) : (_259_v73)), _259_v73, _259_v73, _259_v73);
          _264_v78 = _264_v78;
          let _265_v79;
          let _nw28 = new _module.C3();
          _nw28.__ctor(!((_233_v55).f19) || ((_233_v55).f19), ((_163_v0) ? ((_233_v55).f19) : ((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])));
          _265_v79 = _nw28;
          let _266_v80;
          let _nw29 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _266_v80 = _nw29;
          let _nw30 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _266_v80 = _nw30;
          _166_v3 = _166_v3;
          let _267_v81;
          let _nw31 = new _module.C3();
          _nw31.__ctor((_233_v55).f19, (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))]);
          _267_v81 = _nw31;
        } else {
          let _268_v82;
          let _out3;
          _out3 = (_233_v55).m1((_188_v20)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length))], _170_globalState);
          _268_v82 = _out3;
          let _269_v83;
          _269_v83 = _module.D1.create_DC5(_169_v6, _166_v3, _165_v2);
          let _270_v84;
          let _271_v85;
          let _272_v86;
          let _273_v87;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector0 = (_233_v55).m2((_188_v20)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length))], _dafny.Seq.update(_169_v6, _module.__default.safeIndex(new BigNumber(-928), new BigNumber((_169_v6).length)), _175_v10), (_269_v83).dtor_cf10, _170_globalState);
          _out4 = _outcollector0[0];
          _out5 = _outcollector0[1];
          _out6 = _outcollector0[2];
          _out7 = _outcollector0[3];
          _270_v84 = _out4;
          _271_v85 = _out5;
          _272_v86 = _out6;
          _273_v87 = _out7;
          let _index47 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length));
          (_188_v20)[_index47] = _module.__default.safeDivisionInt(((_259_v73)[_module.__default.safeIndex((_188_v20)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length))], new BigNumber((_259_v73).length))]).plus(_165_v2), (_dafny.ZERO).minus((_188_v20)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length))]));
          let _index48 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length));
          (_188_v20)[_index48] = (_module.__default.fm31(_170_globalState)).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_174_v9).update(_165_v2, _module.__default.abs((_188_v20)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_188_v20).length))]))).cardinality()), _165_v2)));
          _166_v3 = _166_v3;
        }
        let _index49 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        (_166_v3)[_index49] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("lv"), _169_v6);
      } else if (_source7.is_DC14) {
        let _274___mcc_h0 = (_source7).cf23;
        let _275_cf23 = _274___mcc_h0;
        let _276_v88;
        let _nw32 = Array((new BigNumber(3)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _166_v3;
        _nw32[(_dafny.ONE).toNumber()] = _166_v3;
        _nw32[(new BigNumber(2)).toNumber()] = _166_v3;
        _276_v88 = _nw32;
        let _index50 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_276_v88).length));
        (_276_v88)[_index50] = ((_275_cf23) ? (_166_v3) : (_166_v3));
        let _index51 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_276_v88).length));
        (_276_v88)[_index51] = _166_v3;
        let _277_v89;
        _277_v89 = _dafny.Seq.of(_165_v2);
        let _278_v90;
        _278_v90 = _dafny.Seq.of(_166_v3, _166_v3);
        let _279_v91;
        _279_v91 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_277_v89, _module.__default.fm16(new BigNumber((_278_v90).length), _165_v2, (_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))], _165_v2, _170_globalState)),new BigNumber((_169_v6).length));
        _279_v91 = (_279_v91).update(_277_v89, _module.__default.safeModuloInt(_165_v2, _165_v2));
        let _280_v92;
        _280_v92 = _dafny.MultiSet.fromElements(_192_v22, _192_v22, _192_v22);
        let _pat_let_tv9 = _173_v8;
        _169_v6 = _module.__default.fm17((_dafny.MultiSet.fromElements(_192_v22, _192_v22, function (_pat_let4_0) {
          return function (_281_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_282_dt__update_hcf3_h1) {
                return _module.D1.create_DC2(_282_dt__update_hcf3_h1);
              }(_pat_let5_0);
            }(_pat_let_tv9);
          }(_pat_let4_0);
        }(_192_v22))).Intersect(_280_v92), _170_globalState);
        _165_v2 = (new BigNumber(297)).minus(_165_v2);
      } else if (_source7.is_DC12) {
        let _283___mcc_h1 = (_source7).cf22;
        let _284_cf22 = _283___mcc_h1;
        let _285_v93;
        let _nw33 = new _module.C2();
        _nw33.__ctor(!(!(_163_v0)));
        _285_v93 = _nw33;
        let _index52 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
        (_249_v66)[_index52] = _188_v20;
        let _index53 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
        (_249_v66)[_index53] = _188_v20;
        let _286_v94;
        let _nw34 = new _module.C5();
        _nw34.__ctor(!(true) || (_163_v0), _188_v20);
        _286_v94 = _nw34;
      } else {
        let _287___mcc_h2 = (_source7).cf24;
        let _288_cf24 = _287___mcc_h2;
        if ((_233_v55).f19) {
          _163_v0 = (_233_v55).f19;
          _163_v0 = (_233_v55).f19;
          let _index54 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
          (_249_v66)[_index54] = _188_v20;
          let _289_v95;
          let _out8;
          _out8 = (_233_v55).m1(_165_v2, _170_globalState);
          _289_v95 = _out8;
          _163_v0 = (_233_v55).f19;
        } else {
          let _290_v96;
          let _291_v97;
          let _292_v98;
          let _293_v99;
          let _out9;
          let _out10;
          let _out11;
          let _out12;
          let _outcollector1 = (_233_v55).m2(_165_v2, _169_v6, (_dafny.ZERO).minus(_165_v2), _170_globalState);
          _out9 = _outcollector1[0];
          _out10 = _outcollector1[1];
          _out11 = _outcollector1[2];
          _out12 = _outcollector1[3];
          _290_v96 = _out9;
          _291_v97 = _out10;
          _292_v98 = _out11;
          _293_v99 = _out12;
          let _294_v100;
          let _nw35 = new _module.C5();
          _nw35.__ctor(_291_v97, _188_v20);
          _294_v100 = _nw35;
          let _295_v101;
          _295_v101 = _dafny.Map.Empty.slice().updateUnsafe(_165_v2,_294_v100);
          _292_v98 = !(!((_165_v2).isLessThanOrEqualTo(((_dafny.ZERO).minus(_165_v2)).multipliedBy(new BigNumber((_295_v101).length)))));
          let _296_v102;
          _296_v102 = _dafny.Seq.of(_166_v3);
          let _297_v103;
          _297_v103 = _module.D1.create_DC4((_296_v102)[_module.__default.safeIndex(_165_v2, new BigNumber((_296_v102).length))], _165_v2, _175_v10);
          _292_v98 = ((_297_v103).dtor_cf6).isLessThan(_165_v2);
          _165_v2 = _module.__default.safeDivisionInt((_165_v2).plus((_dafny.ZERO).minus(_165_v2)), _165_v2);
          let _index55 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_188_v20).length));
          (_188_v20)[_index55] = (_165_v2).minus(_165_v2);
          let _index56 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_188_v20).length));
          (_188_v20)[_index56] = _165_v2;
        }
        if (_163_v0) {
          let _298_v104;
          _298_v104 = _dafny.Seq.of(_165_v2, _165_v2, new BigNumber((_174_v9).cardinality()));
          let _299_v105;
          _299_v105 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("epfsk")).length), _165_v2), _298_v104, _298_v104);
          let _300_v106;
          let _301_v107;
          let _302_v108;
          let _303_v109;
          let _out13;
          let _out14;
          let _out15;
          let _out16;
          let _outcollector2 = (_233_v55).m2((new BigNumber((_dafny.Set.fromElements((_233_v55).f19, (_233_v55).f19)).length)).plus(new BigNumber(((_299_v105)[_module.__default.safeIndex(_165_v2, new BigNumber((_299_v105).length))]).length)), _169_v6, _165_v2, _170_globalState);
          _out13 = _outcollector2[0];
          _out14 = _outcollector2[1];
          _out15 = _outcollector2[2];
          _out16 = _outcollector2[3];
          _300_v106 = _out13;
          _301_v107 = _out14;
          _302_v108 = _out15;
          _303_v109 = _out16;
          let _arr0 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _index57 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          _arr0[_index57] = _module.__default.safeModuloInt(_165_v2, new BigNumber((_303_v109).length));
          let _304_v110;
          let _init3 = ((_305_v108, _306_v107, _307_v3, _308_v0, _309_v55) => function (_310_i13) {
            return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_305_v108,_306_v107), _dafny.Map.Empty.slice().updateUnsafe((_307_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_307_v3).length))],_308_v0))).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,_305_v108), _dafny.Map.Empty.slice().updateUnsafe(_306_v107,(_309_v55).f19)));
          })(_302_v108, _301_v107, _166_v3, _163_v0, _233_v55);
          let _nw36 = Array((new BigNumber(23)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw36.length); _i0_3++) {
            _nw36[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _304_v110 = _nw36;
          let _311_v111;
          _311_v111 = _dafny.Map.Empty.slice().updateUnsafe(_163_v0,!(_301_v107));
          let _312_v112;
          _312_v112 = _dafny.Set.fromElements(_311_v111, _311_v111);
          let _index58 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_304_v110).length));
          (_304_v110)[_index58] = _312_v112;
          let _313_v113;
          _313_v113 = _dafny.Set.fromElements(_165_v2);
          let _arr1 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _index59 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_304_v110).length));
          let _rhs31 = (_165_v2).plus(new BigNumber((_313_v113).length));
          let _rhs32 = _165_v2;
          let _rhs33 = _213_v39;
          let _rhs34 = _165_v2;
          let _rhs35 = _312_v112;
          let _lhs24 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _lhs25 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          let _lhs26 = _304_v110;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_304_v110).length));
          _165_v2 = _rhs31;
          _165_v2 = _rhs32;
          _213_v39 = _rhs33;
          _lhs24[_lhs25] = _rhs34;
          _lhs26[_lhs27] = _rhs35;
          let _arr2 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _index61 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          let _rhs36 = ((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))])[_module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length))];
          let _rhs37 = _dafny.Map.Empty.slice().updateUnsafe(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))])[_module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length))],_165_v2);
          let _lhs28 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _lhs29 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          _lhs28[_lhs29] = _rhs36;
          _173_v8 = _rhs37;
          let _314_v114;
          _314_v114 = _dafny.Map.Empty.slice().updateUnsafe(true,_298_v104);
          let _arr3 = (_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))];
          let _index62 = _module.__default.safeIndex(new BigNumber(267), new BigNumber(((_249_v66)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length))]).length));
          _arr3[_index62] = _module.__default.safeModuloInt(new BigNumber((_314_v114).length), _module.__default.fm13(_302_v108, (_233_v55).f19, _175_v10, _170_globalState));
          let _index63 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          (_166_v3)[_index63] = ((_233_v55).f19) === ((_233_v55).f19);
        } else {
          let _index64 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          (_166_v3)[_index64] = false;
          let _315_v116;
          _315_v116 = _dafny.Seq.of(new BigNumber((_169_v6).length), _165_v2);
          let _index65 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
          (_166_v3)[_index65] = _module.__default.fm0(function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_315_v116).Elements) {
              let _316_v115 = _compr_26;
              if (_dafny.Seq.contains(_315_v116, _316_v115)) {
                _coll26.push([(_316_v115).plus(_165_v2),new BigNumber(((_174_v9).update(new BigNumber(272), _module.__default.abs(_165_v2))).cardinality())]);
              }
            }
            return _coll26;
          }(), _165_v2, _170_globalState);
          let _317_v117;
          _317_v117 = _module.D2.create_DC6(_315_v116);
          let _318_v118;
          _318_v118 = _module.D2.create_DC6((_317_v117).dtor_cf11);
          let _319_v119;
          _319_v119 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_165_v2),_169_v6);
          let _320_v120;
          _320_v120 = _dafny.Map.Empty.slice().updateUnsafe((_233_v55).f19,_319_v119);
          let _rhs38 = new BigNumber((_dafny.Seq.UnicodeFromString("od")).length);
          let _rhs39 = _module.__default.fm48(_175_v10, ((((_320_v120).contains((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])) ? ((_320_v120).get((_166_v3)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length))])) : (_319_v119))).update(_165_v2, _169_v6), _165_v2, _dafny.Seq.IsProperPrefixOf(_169_v6, _169_v6), _170_globalState);
          let _rhs40 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), ((_321_v2) => function (_322_i14) {
            return _321_v2;
          })(_165_v2));
          let _rhs41 = (_203_v32).update((_233_v55).f19, new BigNumber(488));
          _165_v2 = _rhs38;
          _318_v118 = _rhs39;
          _315_v116 = _rhs40;
          _203_v32 = _rhs41;
          let _index66 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_249_v66).length));
          let _nw37 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          (_249_v66)[_index66] = _nw37;
          let _323_v121;
          let _nw38 = new _module.C5();
          _nw38.__ctor(true, _188_v20);
          _323_v121 = _nw38;
          _323_v121 = _323_v121;
        }
        let _324_v122;
        let _out17;
        _out17 = (_233_v55).m1((new BigNumber(85)).multipliedBy(_165_v2), _170_globalState);
        _324_v122 = _out17;
        let _index67 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        let _rhs42 = _163_v0;
        let _rhs43 = (_233_v55).fm27(_170_globalState);
        let _lhs30 = _166_v3;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_166_v3).length));
        _lhs30[_lhs31] = _rhs42;
        _175_v10 = _rhs43;
      }
      _165_v2 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(_165_v2))).multipliedBy(_165_v2);
      process.stdout.write(_dafny.toString(_163_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_164_v1, _dafny.Seq.of(false, true, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_165_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v3)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v4).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v5).equals(_dafny.Set.fromElements(_dafny.Seq.of(false, true, false, true), _dafny.Seq.of(false, false, true, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_169_v6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f1)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_170_globalState).f4, _dafny.Seq.of(false, true, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f5).equals(_dafny.Set.fromElements(_dafny.Seq.of(false, true, false, true), _dafny.Seq.of(false, false, true, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f7).equals(_dafny.MultiSet.fromElements(new BigNumber(-199)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_170_globalState).f10)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_170_globalState).f11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v7).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v7).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_173_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(199),new BigNumber(833)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v9).equals(_dafny.MultiSet.fromElements(new BigNumber(-199)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_176_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v20)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_191_v21).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_192_v22).dtor_cf3).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(199),new BigNumber(833)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_193_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v32).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(78001)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_211_v37).equals(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v38).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v39).equals(_dafny.MultiSet.fromElements(false, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_214_v41, _dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_216_v42)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_230_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v55).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_234_v56).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v66)[new BigNumber(6)])[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO);
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
    static create_DC3(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC4(cf5, cf6, cf7) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9, cf10) {
      let $dt = new D1(2);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D1(3);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC2() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + this.cf8.toVerbatimString(true) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO);
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
    static create_DC7(cf12, cf13, cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf16, cf17, cf18) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC6(cf11) {
      let $dt = new D2(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + this.cf12.toVerbatimString(true) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf16 === other.cf16 && this.cf17 === other.cf17 && this.cf18 === other.cf18;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC9(cf19) {
      let $dt = new D3(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
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
    static create_DC11(cf21) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC10(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(false);
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
    static create_DC13() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D5(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC12(cf22) {
      let $dt = new D5(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC15(cf24) {
      let $dt = new D5(3);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13();
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
    static create_DC17(cf26, cf27, cf28) {
      let $dt = new D6(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC16(cf25) {
      let $dt = new D6(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC18(cf29) {
      let $dt = new D6(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(_dafny.ZERO, _dafny.ZERO, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf30) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D7(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf30 === other.cf30;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf34) {
      let $dt = new D8(0);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf36, cf37, cf38) {
      let $dt = new D9(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf39, cf40, cf41) {
      let $dt = new D9(1);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf35) {
      let $dt = new D9(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(false, false, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf42) {
      let $dt = new D10(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf44) {
      let $dt = new D11(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC29(cf45, cf46, cf47, cf48) {
      let $dt = new D11(1);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC30(cf49) {
      let $dt = new D11(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC27(cf43) {
      let $dt = new D11(3);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC31(cf50) {
      let $dt = new D11(4);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC27() { return this.$tag === 3; }
    get is_DC31() { return this.$tag === 4; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 4) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf49 === other.cf49;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf52, cf53, cf54, cf55, cf56) {
      let $dt = new D12(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf51) {
      let $dt = new D12(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC34(cf57) {
      let $dt = new D12(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + this.cf52.toVerbatimString(true) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55 && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this._f0 = _dafny.ZERO;
      this._f1 = [];
      this._f2 = false;
      this._f3 = false;
      this._f4 = _dafny.Seq.of();
      this._f5 = _dafny.Set.Empty;
      this._f6 = false;
      this._f7 = _dafny.MultiSet.Empty;
      this._f8 = false;
      this._f9 = _dafny.ZERO;
      this._f10 = [];
      this._f11 = _dafny.Seq.UnicodeFromString("");
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f15) {
      let _this = this;
      (_this)._f15 = f15;
      return;
    }
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length)))).contains(!(false) || (false));
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f14 = [];
      this.f18 = _dafny.Seq.UnicodeFromString("");
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f17, f18, f14) {
      let _this = this;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this)._f14 = f14;
      return;
    }
    fm6(globalState) {
      let _this = this;
      return (_this).f17;
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(-859);
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements((_this).f17)).length), new BigNumber(41))).isLessThanOrEqualTo(new BigNumber((_this.f18).length)));
    };
    fm15(globalState) {
      let _this = this;
      return new BigNumber((((_dafny.MultiSet.fromElements((_this).f17)).Difference(_dafny.MultiSet.fromElements((_this).f17))).Union(_dafny.MultiSet.fromElements(false, (_this).f17))).cardinality());
    };
    m5(p0, globalState) {
      let _this = this;
      let _325_i0;
      _325_i0 = _dafny.ZERO;
      L4: {
        while ((_this).f17) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_325_i0)) {
              break L4;
            }
            _325_i0 = (_325_i0).plus(_dafny.ONE);
            let _326_v0;
            _326_v0 = new BigNumber(669);
            let _327_v1;
            _327_v1 = true;
            let _rhs44 = _module.__default.safeDivisionInt(new BigNumber(703), (_326_v0).multipliedBy(_326_v0));
            let _rhs45 = (_this).f17;
            let _rhs46 = (_326_v0).isLessThanOrEqualTo(_326_v0);
            _326_v0 = _rhs44;
            _327_v1 = _rhs45;
            _327_v1 = _rhs46;
            let _328_v2;
            _328_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm14(_326_v0, (_this).f17, _326_v0, _326_v0, globalState),(_this).f14);
            let _329_v3;
            _329_v3 = _dafny.Map.Empty.slice().updateUnsafe(_326_v0,(((_328_v2).contains((_this).f17)) ? ((_328_v2).get((_this).f17)) : ((_this).f14)));
            let _rhs47 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_326_v0,(_this).f14)).Merge(_329_v3)).length);
            let _rhs48 = false;
            _326_v0 = _rhs47;
            _327_v1 = _rhs48;
            let _330_v4;
            _330_v4 = new _dafny.CodePoint('t'.codePointAt(0));
            _330_v4 = (_this.f18)[_module.__default.safeIndex(_326_v0, new BigNumber((_this.f18).length))];
            let _index68 = _module.__default.safeIndex(new BigNumber(585), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index68] = (_326_v0).minus(_326_v0);
            let _331_v5;
            _331_v5 = _dafny.Set.fromElements(true, _327_v1, (_this).f17);
            let _332_v6;
            _332_v6 = _dafny.Map.Empty.slice().updateUnsafe((_331_v5).Difference(_dafny.Set.fromElements((_this).f17, true)),((_dafny.ZERO).minus(_326_v0)).minus(_326_v0));
            let _index69 = _module.__default.safeIndex(new BigNumber(585), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index69] = new BigNumber((_332_v6).length);
          }
        }
      }
      let _333_v7;
      _333_v7 = new BigNumber(-432);
      let _hi2 = new BigNumber(98);
      for (let _334_i1 = _333_v7; _334_i1.isLessThan(_hi2); _334_i1 = _334_i1.plus(_dafny.ONE)) {
        let _335_v8;
        _335_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_334_i1);
        let _336_v9;
        _336_v9 = _dafny.Seq.of((_this).f17);
        let _337_v10;
        _337_v10 = _dafny.Map.Empty.slice().updateUnsafe(_336_v9,_333_v7);
        let _338_v11;
        _338_v11 = _dafny.Seq.of(_334_i1, new BigNumber(((_337_v10).update(_dafny.Seq.of((_this).f17, (_this).f17), _334_i1)).length), new BigNumber(740));
        let _339_v12;
        _339_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(true),(((_this).f17) ? (_module.__default.fm16(_333_v7, new BigNumber((_335_v8).length), (_this).f17, _334_i1, globalState)) : (_338_v11)));
        _339_v12 = (_339_v12).update(!((_this).f17) || ((_this).f17), _338_v11);
        _333_v7 = (_module.__default.safeDivisionInt(_334_i1, _333_v7)).multipliedBy(_333_v7);
        _333_v7 = _module.__default.safeDivisionInt(new BigNumber(160), new BigNumber(5));
        let _340_v13;
        _340_v13 = _dafny.MultiSet.fromElements(_334_i1);
        let _341_v14;
        _341_v14 = _dafny.Set.fromElements((_this).f17);
        let _342_v15;
        _342_v15 = _dafny.Set.fromElements(_341_v14);
        _333_v7 = (((_340_v13).contains((_333_v7).plus(new BigNumber((_this.f18).length)))) ? ((_340_v13).get((_333_v7).plus(new BigNumber((_this.f18).length)))) : (new BigNumber((_342_v15).length)));
      }
      let _343_i2;
      _343_i2 = _dafny.ZERO;
      L5: {
        while (!(!((_this).f17))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_343_i2)) {
              break L5;
            }
            _343_i2 = (_343_i2).plus(_dafny.ONE);
            let _344_v16;
            let _nw39 = new _module.C0();
            _nw39.__ctor(_333_v7);
            _344_v16 = _nw39;
            let _345_v17;
            _345_v17 = _dafny.Map.Empty.slice().updateUnsafe(_333_v7,(_344_v16).f15);
            let _346_v18;
            _346_v18 = _module.D1.create_DC2(_345_v17);
            let _source8 = _346_v18;
            if (_source8.is_DC3) {
              let _347___mcc_h0 = (_source8).cf4;
              let _348_cf4 = _347___mcc_h0;
              let _349_v19;
              _349_v19 = _module.D1.create_DC3(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), function (_350_i3) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length));
              let _351_v20;
              let _nw40 = Array((new BigNumber(26)).toNumber());
              _nw40[(_dafny.ZERO).toNumber()] = _module.D1.create_DC3(new BigNumber((_this.f18).length));
              _nw40[(_dafny.ONE).toNumber()] = _349_v19;
              _nw40[(new BigNumber(2)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(3)).toNumber()] = _module.D1.create_DC3((_344_v16).f15);
              _nw40[(new BigNumber(4)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(5)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(6)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(7)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(8)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(9)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(10)).toNumber()] = _module.D1.create_DC3(_348_cf4);
              _nw40[(new BigNumber(11)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(12)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(13)).toNumber()] = _module.D1.create_DC3(new BigNumber(115));
              _nw40[(new BigNumber(14)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(15)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(16)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(17)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(18)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(19)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(20)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(21)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(22)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(23)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(24)).toNumber()] = _349_v19;
              _nw40[(new BigNumber(25)).toNumber()] = _349_v19;
              _351_v20 = _nw40;
              let _index70 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_351_v20).length));
              (_351_v20)[_index70] = _349_v19;
              let _352_v21;
              _352_v21 = true;
              let _353_v22;
              _353_v22 = _dafny.Seq.of(new BigNumber((_this.f18).length), new BigNumber(31), _348_cf4);
              let _pat_let_tv10 = _353_v22;
              let _pat_let_tv11 = _348_cf4;
              let _pat_let_tv12 = _352_v21;
              let _index71 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_351_v20).length));
              let _rhs49 = function (_pat_let6_0) {
                return function (_354_dt__update__tmp_h0) {
                  return function (_pat_let7_0) {
                    return function (_355_dt__update_hcf4_h0) {
                      return _module.D1.create_DC3(_355_dt__update_hcf4_h0);
                    }(_pat_let7_0);
                  }(((_pat_let_tv12) ? (new BigNumber((_pat_let_tv10).length)) : (_pat_let_tv11)));
                }(_pat_let6_0);
              }(_module.D1.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("x")).length)));
              let _rhs50 = _module.__default.safeDivisionInt(_333_v7, _348_cf4);
              let _rhs51 = !((_this).f17);
              let _rhs52 = (_344_v16).f15;
              let _rhs53 = (_this).f17;
              let _lhs32 = _351_v20;
              let _lhs33 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_351_v20).length));
              _lhs32[_lhs33] = _rhs49;
              _333_v7 = _rhs50;
              _352_v21 = _rhs51;
              _348_cf4 = _rhs52;
              _352_v21 = _rhs53;
              let _356_v23;
              _356_v23 = _module.D2.create_DC7(_this.f18, (_344_v16).f15, (_344_v16).f15, (_this).f17);
              let _357_v24;
              _357_v24 = _dafny.Seq.of(_352_v21, _352_v21);
              let _358_v25;
              _358_v25 = new _dafny.CodePoint('l'.codePointAt(0));
              let _pat_let_tv13 = _345_v17;
              let _359_v26;
              _359_v26 = _dafny.MultiSet.fromElements(_346_v18, function (_pat_let8_0) {
                return function (_360_dt__update__tmp_h1) {
                  return function (_pat_let9_0) {
                    return function (_361_dt__update_hcf3_h0) {
                      return _module.D1.create_DC2(_361_dt__update_hcf3_h0);
                    }(_pat_let9_0);
                  }(_pat_let_tv13);
                }(_pat_let8_0);
              }(_346_v18));
              let _362_v27;
              let _nw41 = Array((new BigNumber(25)).toNumber());
              _nw41[(_dafny.ZERO).toNumber()] = _352_v21;
              _nw41[(_dafny.ONE).toNumber()] = (_356_v23).dtor_cf15;
              _nw41[(new BigNumber(2)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(3)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(4)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(5)).toNumber()] = (_333_v7).isLessThan(new BigNumber((_357_v24).length));
              _nw41[(new BigNumber(6)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(7)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(8)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(9)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(10)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(11)).toNumber()] = !(!(_352_v21));
              _nw41[(new BigNumber(12)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("eyusi"), _358_v25);
              _nw41[(new BigNumber(13)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(14)).toNumber()] = (_357_v24)[_module.__default.safeIndex(new BigNumber((_353_v22).length), new BigNumber((_357_v24).length))];
              _nw41[(new BigNumber(15)).toNumber()] = !(_352_v21);
              _nw41[(new BigNumber(16)).toNumber()] = false;
              _nw41[(new BigNumber(17)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(18)).toNumber()] = false;
              _nw41[(new BigNumber(19)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(20)).toNumber()] = (_352_v21) === ((_this).f17);
              _nw41[(new BigNumber(21)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(22)).toNumber()] = _352_v21;
              _nw41[(new BigNumber(23)).toNumber()] = (_this).f17;
              _nw41[(new BigNumber(24)).toNumber()] = !_dafny.areEqual(_module.__default.fm17(_359_v26, globalState), _this.f18);
              _362_v27 = _nw41;
              let _index72 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_362_v27).length));
              (_362_v27)[_index72] = !(!((_this).f17) || ((_this).f17));
              let _index73 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_362_v27).length));
              (_362_v27)[_index73] = ((_344_v16).f15).isEqualTo(_module.__default.safeModuloInt(new BigNumber(250), new BigNumber(848)));
              let _363_v28;
              _363_v28 = _dafny.Seq.of(_dafny.Seq.of((_344_v16).f15));
              let _364_v29;
              _364_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_344_v16).f15, (_344_v16).f15));
              let _rhs54 = (_363_v28)[_module.__default.safeIndex((_344_v16).f15, new BigNumber((_363_v28).length))];
              let _rhs55 = (_dafny.MultiSet.fromElements(_dafny.Seq.update(_353_v22, _module.__default.safeIndex(_333_v7, new BigNumber((_353_v22).length)), (_344_v16).f15), _module.__default.fm16(new BigNumber(263), (_344_v16).f15, !(!((_this).f17)), _333_v7, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), ((_365_v27) => function (_366_i4) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_365_v27)[_module.__default.safeIndex(new BigNumber(336), new BigNumber((_365_v27).length))],(_this).f17)).length);
              })(_362_v27)))).IsDisjointFrom(_364_v29);
              let _rhs56 = new BigNumber(172);
              _353_v22 = _rhs54;
              _352_v21 = _rhs55;
              _333_v7 = _rhs56;
              let _367_v30;
              let _368_v31;
              let _369_v32;
              let _out18;
              let _out19;
              let _out20;
              let _outcollector3 = (_this).m6((_362_v27)[_module.__default.safeIndex(new BigNumber(336), new BigNumber((_362_v27).length))], globalState);
              _out18 = _outcollector3[0];
              _out19 = _outcollector3[1];
              _out20 = _outcollector3[2];
              _367_v30 = _out18;
              _368_v31 = _out19;
              _369_v32 = _out20;
            } else if (_source8.is_DC4) {
              let _370___mcc_h1 = (_source8).cf5;
              let _371___mcc_h2 = (_source8).cf6;
              let _372___mcc_h3 = (_source8).cf7;
              let _373_cf7 = _372___mcc_h3;
              let _374_cf6 = _371___mcc_h2;
              let _375_cf5 = _370___mcc_h1;
              let _376_v33;
              _376_v33 = _dafny.Set.fromElements((_this).f17, (_this).f17);
              let _377_v34;
              _377_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f17) ? (_this.f18) : (_dafny.Seq.UnicodeFromString("mkrgtsv"))),_module.__default.safeDivisionInt(_333_v7, new BigNumber((_376_v33).length)));
              _377_v34 = (_377_v34).update(_this.f18, (_344_v16).f15);
              let _378_v35;
              _378_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
              _374_cf6 = (_dafny.ZERO).minus((_333_v7).multipliedBy(new BigNumber((_378_v35).length)));
              (_this).f18 = _dafny.Seq.update(_dafny.Seq.Concat(_this.f18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), ((_379_cf7) => function (_380_i5) {
                return _379_cf7;
              })(_373_cf7))), _module.__default.safeIndex((_344_v16).f15, new BigNumber((_dafny.Seq.Concat(_this.f18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), ((_381_cf7) => function (_382_i5) {
                return _381_cf7;
              })(_373_cf7)))).length)), _373_cf7);
              let _383_v36;
              _383_v36 = true;
              _383_v36 = ((new BigNumber(-718)).plus(_374_cf6)).isLessThanOrEqualTo(new BigNumber(447));
            } else if (_source8.is_DC5) {
              let _384___mcc_h4 = (_source8).cf8;
              let _385___mcc_h5 = (_source8).cf9;
              let _386___mcc_h6 = (_source8).cf10;
              let _387_cf10 = _386___mcc_h6;
              let _388_cf9 = _385___mcc_h5;
              let _389_cf8 = _384___mcc_h4;
              let _390_v37;
              _390_v37 = _dafny.Map.Empty.slice().updateUnsafe(_333_v7,(_this).f17);
              let _391_v38;
              let _nw42 = Array((new BigNumber(21)).toNumber());
              _nw42[(_dafny.ZERO).toNumber()] = _346_v18;
              _nw42[(_dafny.ONE).toNumber()] = _module.D1.create_DC2(_345_v17);
              _nw42[(new BigNumber(2)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(3)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(4)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(5)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(6)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(7)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(8)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(9)).toNumber()] = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm18((_this).f17, (_this).f17, (_344_v16).f15, new BigNumber(348), globalState)).length),new BigNumber((_390_v37).length)));
              _nw42[(new BigNumber(10)).toNumber()] = ((!(true)) ? (_346_v18) : (_346_v18));
              _nw42[(new BigNumber(11)).toNumber()] = _module.D1.create_DC2(_345_v17);
              _nw42[(new BigNumber(12)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(13)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(14)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(15)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(16)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(17)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(18)).toNumber()] = _346_v18;
              _nw42[(new BigNumber(19)).toNumber()] = _module.D1.create_DC2(_345_v17);
              _nw42[(new BigNumber(20)).toNumber()] = _module.__default.fm19(globalState);
              _391_v38 = _nw42;
              let _index74 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_391_v38).length));
              (_391_v38)[_index74] = _346_v18;
              let _index75 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_391_v38).length));
              (_391_v38)[_index75] = _346_v18;
              let _392_v39;
              let _nw43 = new _module.C0();
              _nw43.__ctor(_333_v7);
              _392_v39 = _nw43;
              let _393_v40;
              _393_v40 = false;
              let _394_v41;
              _394_v41 = _dafny.Seq.of(_393_v40);
              _393_v40 = (false) || ((_394_v41)[_module.__default.safeIndex(_387_cf10, new BigNumber((_394_v41).length))]);
              let _395_v42;
              _395_v42 = _dafny.Seq.of((_dafny.ZERO).minus((_344_v16).f15), (_344_v16).f15);
              _333_v7 = new BigNumber((_dafny.Seq.Concat(_395_v42, _dafny.Seq.Concat(_395_v42, _395_v42))).length);
            } else {
              let _396___mcc_h7 = (_source8).cf3;
              let _397_cf3 = _396___mcc_h7;
              let _index76 = _module.__default.safeIndex(new BigNumber(731), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index76] = (_333_v7).plus((_344_v16).f15);
              let _index77 = _module.__default.safeIndex(new BigNumber(731), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index77] = (_344_v16).f15;
              let _398_v43;
              let _nw44 = Array((new BigNumber(25)).toNumber()).fill(false);
              _398_v43 = _nw44;
              let _index78 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_398_v43).length));
              (_398_v43)[_index78] = (_this).f17;
              let _399_v44;
              _399_v44 = false;
              let _400_v45;
              _400_v45 = _dafny.MultiSet.fromElements(_399_v44, !((_this).f17));
              let _index79 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_398_v43).length));
              let _rhs57 = !(!((new BigNumber((_400_v45).cardinality())).plus((_344_v16).f15)).isEqualTo((_344_v16).f15));
              let _rhs58 = ((_344_v16).f15).isLessThan(new BigNumber((_this.f18).length));
              let _lhs34 = _398_v43;
              let _lhs35 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_398_v43).length));
              _lhs34[_lhs35] = _rhs57;
              _399_v44 = _rhs58;
              let _index80 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_398_v43).length));
              (_398_v43)[_index80] = !(true);
              _399_v44 = (_398_v43)[_module.__default.safeIndex(new BigNumber(993), new BigNumber((_398_v43).length))];
            }
            let _401_v46;
            _401_v46 = _dafny.Seq.of((_this).f17, !((_this).f17));
            let _402_v47;
            _402_v47 = _dafny.Map.Empty.slice().updateUnsafe((_344_v16).f15,_dafny.Seq.Concat(_dafny.Seq.of((_this).f17, (_this).f17), _401_v46));
            let _403_v48;
            _403_v48 = _dafny.Set.fromElements(_401_v46, _dafny.Seq.of((_this).f17, (_this).f17));
            let _404_v49;
            _404_v49 = _dafny.MultiSet.fromElements(true);
            let _405_v50;
            _405_v50 = _dafny.Seq.of(new BigNumber((_403_v48).length), new BigNumber((_404_v49).cardinality()));
            _402_v47 = (_402_v47).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f17, (_this).f17),(_405_v50)[_module.__default.safeIndex(_333_v7, new BigNumber((_405_v50).length))])).length), _401_v46);
            let _406_v51;
            _406_v51 = _module.D0.create_DC0((_this).f17);
            let _source9 = _406_v51;
            if (_source9.is_DC1) {
              let _407___mcc_h8 = (_source9).cf1;
              let _408___mcc_h9 = (_source9).cf2;
              let _409_cf2 = _408___mcc_h9;
              let _410_cf1 = _407___mcc_h8;
              let _index81 = _module.__default.safeIndex(new BigNumber(810), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index81] = _module.__default.safeDivisionInt((_344_v16).f15, _409_cf2);
              let _index82 = _module.__default.safeIndex(new BigNumber(810), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index82] = _409_cf2;
              let _411_v52;
              let _init4 = ((_412_v16) => function (_413_i6) {
                return (_413_i6).multipliedBy((_412_v16).f15);
              })(_344_v16);
              let _nw45 = Array((new BigNumber(19)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw45.length); _i0_4++) {
                _nw45[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _411_v52 = _nw45;
              let _init5 = ((_414_v16) => function (_415_i7) {
                return (_415_i7).multipliedBy((_414_v16).f15);
              })(_344_v16);
              let _nw46 = Array((new BigNumber(19)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw46.length); _i0_5++) {
                _nw46[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _411_v52 = _nw46;
              let _416_v53;
              _416_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_410_cf1);
              let _417_v54;
              _417_v54 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(994));
              let _418_v55;
              _418_v55 = _dafny.Set.fromElements((_344_v16).f15, (_344_v16).f15);
              let _419_v56;
              let _nw47 = Array((new BigNumber(25)).toNumber());
              _nw47[(_dafny.ZERO).toNumber()] = _410_cf1;
              _nw47[(_dafny.ONE).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(2)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(3)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(4)).toNumber()] = !(_410_cf1);
              _nw47[(new BigNumber(5)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(6)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(7)).toNumber()] = false;
              _nw47[(new BigNumber(8)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(9)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(10)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(11)).toNumber()] = false;
              _nw47[(new BigNumber(12)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(13)).toNumber()] = false;
              _nw47[(new BigNumber(14)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(15)).toNumber()] = true;
              _nw47[(new BigNumber(16)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(17)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(18)).toNumber()] = _410_cf1;
              _nw47[(new BigNumber(19)).toNumber()] = true;
              _nw47[(new BigNumber(20)).toNumber()] = (_401_v46)[_module.__default.safeIndex(new BigNumber((_417_v54).length), new BigNumber((_401_v46).length))];
              _nw47[(new BigNumber(21)).toNumber()] = (_this).fm14((_344_v16).f15, _410_cf1, (_344_v16).f15, new BigNumber((_418_v55).length), globalState);
              _nw47[(new BigNumber(22)).toNumber()] = (_this).f17;
              _nw47[(new BigNumber(23)).toNumber()] = false;
              _nw47[(new BigNumber(24)).toNumber()] = _410_cf1;
              _419_v56 = _nw47;
              let _420_v57;
              _420_v57 = _dafny.Seq.of(_416_v53);
              let _421_v58;
              _421_v58 = _dafny.Map.Empty.slice().updateUnsafe(_419_v56,(_420_v57)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-257)), function (_422_i8) {
                return new _dafny.CodePoint('j'.codePointAt(0));
              })).length), new BigNumber((_420_v57).length))]);
              let _423_v59;
              _423_v59 = _416_v53;
              let _424_v60;
              let _nw48 = Array((new BigNumber(10)).toNumber());
              _nw48[(_dafny.ZERO).toNumber()] = _416_v53;
              _nw48[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_410_cf1),!((_this).f17));
              _nw48[(new BigNumber(2)).toNumber()] = (_416_v53).Merge(_416_v53);
              _nw48[(new BigNumber(3)).toNumber()] = (((_421_v58).contains(_419_v56)) ? ((_421_v58).get(_419_v56)) : ((_416_v53).update(_410_cf1, (((_416_v53).contains((_this).f17)) ? ((_416_v53).get((_this).f17)) : (_410_cf1)))));
              _nw48[(new BigNumber(4)).toNumber()] = _416_v53;
              _nw48[(new BigNumber(5)).toNumber()] = (_423_v59);
              _nw48[(new BigNumber(6)).toNumber()] = _416_v53;
              _nw48[(new BigNumber(7)).toNumber()] = _416_v53;
              _nw48[(new BigNumber(8)).toNumber()] = _416_v53;
              _nw48[(new BigNumber(9)).toNumber()] = (_module.__default.fm20(_410_cf1, ((_this).f14)[_module.__default.safeIndex(new BigNumber(810), new BigNumber(((_this).f14).length))], _410_cf1, globalState)).Merge(_416_v53);
              _424_v60 = _nw48;
              let _index83 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_424_v60).length));
              (_424_v60)[_index83] = _416_v53;
              let _index84 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_424_v60).length));
              (_424_v60)[_index84] = _416_v53;
              _411_v52 = _411_v52;
            } else {
              let _425___mcc_h10 = (_source9).cf0;
              let _426_cf0 = _425___mcc_h10;
              let _427_v61;
              let _428_v62;
              let _429_v63;
              let _out21;
              let _out22;
              let _out23;
              let _outcollector4 = (_this).m6((_401_v46)[_module.__default.safeIndex((_344_v16).f15, new BigNumber((_401_v46).length))], globalState);
              _out21 = _outcollector4[0];
              _out22 = _outcollector4[1];
              _out23 = _outcollector4[2];
              _427_v61 = _out21;
              _428_v62 = _out22;
              _429_v63 = _out23;
              _405_v50 = _dafny.Seq.Concat(_405_v50, _405_v50);
              let _430_v64;
              _430_v64 = new _dafny.CodePoint('h'.codePointAt(0));
              let _431_v65;
              _431_v65 = _dafny.Map.Empty.slice().updateUnsafe(_427_v61,_430_v64);
              let _432_v66;
              _432_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_431_v65);
              _432_v66 = (_dafny.Map.Empty.slice().updateUnsafe((_this).f14,_431_v65)).Merge(_432_v66);
              let _433_v67;
              let _init6 = ((_434_v63) => function (_435_i9) {
                return !(_434_v63);
              })(_429_v63);
              let _nw49 = Array((new BigNumber(4)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw49.length); _i0_6++) {
                _nw49[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _433_v67 = _nw49;
              let _436_v68;
              _436_v68 = _dafny.Map.Empty.slice().updateUnsafe(_433_v67,_dafny.Set.fromElements((_this).fm6(globalState)));
              let _437_v69;
              _437_v69 = _dafny.Set.fromElements(_429_v63, true, _429_v63);
              _436_v68 = (_436_v68).update(_433_v67, _437_v69);
            }
          }
        }
      }
      let _438_v70;
      _438_v70 = false;
      _438_v70 = _438_v70;
      let _439_v71;
      let _440_v72;
      let _441_v73;
      let _out24;
      let _out25;
      let _out26;
      let _outcollector5 = (_this).m6((_333_v7).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-78)), function (_442_i10) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length)), globalState);
      _out24 = _outcollector5[0];
      _out25 = _outcollector5[1];
      _out26 = _outcollector5[2];
      _439_v71 = _out24;
      _440_v72 = _out25;
      _441_v73 = _out26;
      let _443_v74;
      let _nw50 = new _module.C0();
      _nw50.__ctor((_this).fm15(globalState));
      _443_v74 = _nw50;
      return;
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _444_v0;
      _444_v0 = new BigNumber(848);
      let _445_v1;
      _445_v1 = _dafny.Map.Empty.slice().updateUnsafe(_444_v0,p0);
      let _446_v2;
      _446_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_445_v1);
      let _447_v4;
      _447_v4 = _dafny.Map.Empty.slice().updateUnsafe(_444_v0,new BigNumber(((((_446_v2).contains((_this).f17)) ? ((_446_v2).get((_this).f17)) : (function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of (_445_v1).Keys.Elements) {
          let _448_v3 = _compr_27;
          if ((_445_v1).contains(_448_v3)) {
            _coll27.push([(_448_v3).plus(new BigNumber(-153)),(_this).f17]);
          }
        }
        return _coll27;
      }()))).length));
      let _449_v5;
      _449_v5 = _module.D0.create_DC1(p0, _444_v0);
      if (_module.__default.fm0(_447_v4, new BigNumber((_dafny.MultiSet.fromElements((_449_v5).dtor_cf1)).cardinality()), globalState)) {
        _444_v0 = _444_v0;
        let _450_v6;
        _450_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,p0);
        let _451_v7;
        _451_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_450_v6);
        let _452_v8;
        let _out27;
        _out27 = _module.__default.m0((((_451_v7).contains(_this.f18)) ? ((_451_v7).get(_this.f18)) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(498)), function (_453_i0) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }),(_this).f17))), globalState);
        _452_v8 = _out27;
        let _rhs59 = _452_v8;
        let _rhs60 = p0;
        r1 = _rhs59;
        r2 = _rhs60;
        let _454_v9;
        let _nw51 = Array((new BigNumber(29)).toNumber()).fill(false);
        _454_v9 = _nw51;
        let _455_v10;
        _455_v10 = _dafny.MultiSet.fromElements(_454_v9);
        r2 = !(((p0) ? (_module.__default.fm0(_447_v4, (_this).fm7(new BigNumber(107), new BigNumber((_455_v10).cardinality()), (_this).f17, globalState), globalState)) : ((_this).f17)));
        let _456_v11;
        let _nw52 = new _module.C0();
        _nw52.__ctor(_444_v0);
        _456_v11 = _nw52;
      } else {
        let _457_v12;
        let _nw53 = Array((new BigNumber(26)).toNumber()).fill(false);
        _457_v12 = _nw53;
        let _index85 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        (_457_v12)[_index85] = p0;
        let _458_v13;
        _458_v13 = _module.D4.create_DC10((_this).f14);
        let _459_v14;
        let _nw54 = Array((new BigNumber(17)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw54[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(2)).toNumber()] = (_458_v13).dtor_cf20;
        _nw54[(new BigNumber(3)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(4)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(6)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(11)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(13)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(14)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(15)).toNumber()] = (_this).f14;
        _nw54[(new BigNumber(16)).toNumber()] = (_this).f14;
        _459_v14 = _nw54;
        let _460_v15;
        let _nw55 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _460_v15 = _nw55;
        let _461_v16;
        _461_v16 = _dafny.Seq.of((_dafny.ZERO).minus(_444_v0), _444_v0, _444_v0);
        let _index86 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_460_v15).length));
        (_460_v15)[_index86] = _dafny.Seq.Concat(_461_v16, _461_v16);
        let _462_v17;
        let _nw56 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _462_v17 = _nw56;
        let _index87 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length));
        (_462_v17)[_index87] = _dafny.Seq.UnicodeFromString("okriy");
        let _index88 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        let _index89 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_460_v15).length));
        let _index90 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length));
        let _rhs61 = false;
        let _rhs62 = _459_v14;
        let _rhs63 = (_444_v0).multipliedBy(_444_v0);
        let _rhs64 = _461_v16;
        let _rhs65 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_463_i1) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }), _this.f18);
        let _lhs36 = _457_v12;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        let _lhs38 = _460_v15;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_460_v15).length));
        let _lhs40 = _462_v17;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length));
        _lhs36[_lhs37] = _rhs61;
        _459_v14 = _rhs62;
        r1 = _rhs63;
        _lhs38[_lhs39] = _rhs64;
        _lhs40[_lhs41] = _rhs65;
        let _464_v18;
        _464_v18 = _module.D1.create_DC3(_444_v0);
        let _pat_let_tv14 = _444_v0;
        let _source10 = function (_pat_let10_0) {
          return function (_465_dt__update__tmp_h0) {
            return function (_pat_let11_0) {
              return function (_466_dt__update_hcf4_h0) {
                return _module.D1.create_DC3(_466_dt__update_hcf4_h0);
              }(_pat_let11_0);
            }(_pat_let_tv14);
          }(_pat_let10_0);
        }(_464_v18);
        if (_source10.is_DC3) {
          let _467___mcc_h0 = (_source10).cf4;
          let _468_cf4 = _467___mcc_h0;
          let _469_v19;
          let _nw57 = new _module.C0();
          _nw57.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus(_444_v0), _468_cf4));
          _469_v19 = _nw57;
          let _470_v20;
          _470_v20 = _module.D1.create_DC2(_447_v4);
          _470_v20 = _470_v20;
          _468_cf4 = (_469_v19).f15;
          _457_v12 = _457_v12;
        } else if (_source10.is_DC4) {
          let _471___mcc_h1 = (_source10).cf5;
          let _472___mcc_h2 = (_source10).cf6;
          let _473___mcc_h3 = (_source10).cf7;
          let _474_cf7 = _473___mcc_h3;
          let _475_cf6 = _472___mcc_h2;
          let _476_cf5 = _471___mcc_h1;
          let _477_v21;
          _477_v21 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).fm7(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qqslcms"), _module.__default.safeIndex(_475_cf6, new BigNumber((_dafny.Seq.UnicodeFromString("qqslcms")).length)), _474_cf7)).length), _444_v0, (_457_v12)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length))], globalState)), _475_cf6, _475_cf6);
          let _478_v22;
          _478_v22 = _dafny.Seq.of(_477_v21, _477_v21, _477_v21, _477_v21, _477_v21);
          let _479_v23;
          _479_v23 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
          let _480_v24;
          _480_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm14(_475_cf6, (_457_v12)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length))], new BigNumber((_479_v23).length), _475_cf6, globalState),_module.__default.fm21(!(p0), globalState));
          _478_v22 = (((_480_v24).contains((_457_v12)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length))])) ? ((_480_v24).get((_457_v12)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length))])) : (_478_v22));
          let _481_v25;
          _481_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_444_v0);
          let _482_v26;
          _482_v26 = _module.D1.create_DC2((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_474_cf7, _474_cf7)).cardinality()),new BigNumber((_481_v25).length))).update(new BigNumber((_dafny.Seq.of(_444_v0, new BigNumber(113))).length), new BigNumber(((_462_v17)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length))]).length)));
          let _483_v27;
          _483_v27 = _dafny.MultiSet.fromElements(_482_v26);
          let _index91 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length));
          (_462_v17)[_index91] = _module.__default.fm17((_483_v27).Union(_dafny.MultiSet.fromElements(_482_v26, _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_475_cf6,_444_v0)), _482_v26, _482_v26, _482_v26)), globalState);
          let _index92 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
          (_457_v12)[_index92] = p0;
          let _484_v28;
          _484_v28 = _dafny.Map.Empty.slice().updateUnsafe(_445_v1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(365)), ((_485_v1, _486_p0, _487_v12) => function (_488_i2) {
            return (_485_v1).update(new BigNumber((_dafny.Seq.of((_this).f17, _486_p0, (_this).f17)).length), (_487_v12)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_487_v12).length))]);
          })(_445_v1, p0, _457_v12)));
          let _489_v29;
          _489_v29 = _dafny.Seq.of(_445_v1, _445_v1);
          let _490_v30;
          _490_v30 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual((((_484_v28).contains(_445_v1)) ? ((_484_v28).get(_445_v1)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), ((_491_v1) => function (_492_i3) {
            return _491_v1;
          })(_445_v1)))), _489_v29),_449_v5);
          _490_v30 = _490_v30;
        } else if (_source10.is_DC5) {
          let _493___mcc_h4 = (_source10).cf8;
          let _494___mcc_h5 = (_source10).cf9;
          let _495___mcc_h6 = (_source10).cf10;
          let _496_cf10 = _495___mcc_h6;
          let _497_cf9 = _494___mcc_h5;
          let _498_cf8 = _493___mcc_h4;
          let _499_v31;
          let _nw58 = new _module.C0();
          _nw58.__ctor(_444_v0);
          _499_v31 = _nw58;
          _499_v31 = _499_v31;
          let _500_v32;
          let _nw59 = new _module.C0();
          _nw59.__ctor(new BigNumber(523));
          _500_v32 = _nw59;
          let _501_v33;
          let _nw60 = new _module.C0();
          _nw60.__ctor((_499_v31).f15);
          _501_v33 = _nw60;
          r0 = _496_cf10;
        } else {
          let _502___mcc_h7 = (_source10).cf3;
          let _503_cf3 = _502___mcc_h7;
          _444_v0 = _444_v0;
          _444_v0 = (((p0) ? (_444_v0) : (_444_v0))).minus(_444_v0);
          let _504_v34;
          _504_v34 = new _dafny.CodePoint('b'.codePointAt(0));
          let _505_v35;
          _505_v35 = _dafny.Map.Empty.slice().updateUnsafe(_444_v0,_504_v34);
          let _506_v36;
          _506_v36 = _dafny.Map.Empty.slice().updateUnsafe(true,_505_v35);
          _506_v36 = (_506_v36).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_505_v35));
          let _507_v37;
          let _nw61 = Array((new BigNumber(5)).toNumber()).fill([]);
          _507_v37 = _nw61;
          let _index93 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_507_v37).length));
          (_507_v37)[_index93] = _462_v17;
          let _508_v38;
          _508_v38 = _dafny.MultiSet.fromElements(p0);
          let _509_v39;
          _509_v39 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(221)), ((_510_v34) => function (_511_i4) {
            return _510_v34;
          })(_504_v34)));
          let _512_v40;
          _512_v40 = _module.D2.create_DC7((_462_v17)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length))], _444_v0, (((_509_v39).contains((_462_v17)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length))])) ? ((_509_v39).get((_462_v17)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_462_v17).length))])) : (new BigNumber(204))), (_this).f17);
          let _513_v41;
          _513_v41 = _dafny.Map.Empty.slice().updateUnsafe(_504_v34,_447_v4);
          let _514_v42;
          _514_v42 = _module.D1.create_DC2((((_513_v41).contains(_504_v34)) ? ((_513_v41).get(_504_v34)) : (_dafny.Map.Empty.slice().updateUnsafe(_444_v0,new BigNumber((_461_v16).length)))));
          let _515_v43;
          _515_v43 = _dafny.MultiSet.fromElements(_514_v42, _514_v42, _module.__default.fm19(globalState));
          let _516_v44;
          _516_v44 = _module.D1.create_DC5(_module.__default.fm17(_515_v43, globalState), _457_v12, _444_v0);
          let _517_v45;
          _517_v45 = _dafny.Set.fromElements(false);
          let _518_v46;
          _518_v46 = _dafny.Seq.of(_517_v45, _517_v45);
          let _index94 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_507_v37).length));
          let _rhs66 = _462_v17;
          let _rhs67 = (_512_v40).dtor_cf14;
          let _rhs68 = (_516_v44).dtor_cf10;
          let _rhs69 = (_508_v38).update((_444_v0).isLessThan(new BigNumber(((_518_v46)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_518_v46).length))]).length)), _module.__default.abs(_444_v0));
          let _rhs70 = _444_v0;
          let _lhs42 = _507_v37;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_507_v37).length));
          _lhs42[_lhs43] = _rhs66;
          r1 = _rhs67;
          r1 = _rhs68;
          _508_v38 = _rhs69;
          _444_v0 = _rhs70;
        }
        _457_v12 = _457_v12;
        let _519_v47;
        _519_v47 = _dafny.Seq.of(_457_v12, _457_v12);
        let _520_v48;
        _520_v48 = new _dafny.CodePoint('m'.codePointAt(0));
        let _521_v49;
        let _nw62 = Array((new BigNumber(19)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = (_519_v47)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_520_v48)).length), new BigNumber((_519_v47).length))];
        _nw62[(_dafny.ONE).toNumber()] = _457_v12;
        _nw62[(new BigNumber(2)).toNumber()] = ((p0) ? (_457_v12) : (_457_v12));
        _nw62[(new BigNumber(3)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(4)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(5)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(6)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(7)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(8)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(9)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(10)).toNumber()] = (_519_v47)[_module.__default.safeIndex(_444_v0, new BigNumber((_519_v47).length))];
        _nw62[(new BigNumber(11)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(12)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(13)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(14)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(15)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(16)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(17)).toNumber()] = _457_v12;
        _nw62[(new BigNumber(18)).toNumber()] = _457_v12;
        _521_v49 = _nw62;
        let _index95 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_521_v49).length));
        (_521_v49)[_index95] = _457_v12;
        let _522_v50;
        _522_v50 = _dafny.Map.Empty.slice().updateUnsafe(!((((_this).f17) ? (p0) : ((_this).f17))),(_this).f17);
        let _index96 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_521_v49).length));
        let _index97 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        let _rhs71 = (_519_v47)[_module.__default.safeIndex(_444_v0, new BigNumber((_519_v47).length))];
        let _rhs72 = (((_522_v50).contains((p0) && ((_this).f17))) ? ((_522_v50).get((p0) && ((_this).f17))) : (p0));
        let _lhs44 = _521_v49;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_521_v49).length));
        let _lhs46 = _457_v12;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        _lhs44[_lhs45] = _rhs71;
        _lhs46[_lhs47] = _rhs72;
        let _index98 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_457_v12).length));
        (_457_v12)[_index98] = (_444_v0).isLessThan(_444_v0);
      }
      let _523_i5;
      _523_i5 = _dafny.ZERO;
      L6: {
        while (p0) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_523_i5)) {
              break L6;
            }
            _523_i5 = (_523_i5).plus(_dafny.ONE);
            r2 = (_this).f17;
            r2 = (_this).f17;
            (_this).f18 = _dafny.Seq.Concat(((false) ? (_this.f18) : (_this.f18)), _this.f18);
            let _524_v51;
            _524_v51 = _dafny.Seq.of(_444_v0, _444_v0, ((false) ? (_444_v0) : ((_dafny.ZERO).minus(_444_v0))));
            let _rhs73 = (_444_v0).minus(_444_v0);
            let _rhs74 = _444_v0;
            let _rhs75 = (_this).f17;
            let _rhs76 = _dafny.Seq.Concat(_524_v51, _524_v51);
            _444_v0 = _rhs73;
            r0 = _rhs74;
            r2 = _rhs75;
            _524_v51 = _rhs76;
          }
        }
      }
      let _525_v52;
      let _nw63 = Array((new BigNumber(4)).toNumber());
      _nw63[(_dafny.ZERO).toNumber()] = p0;
      _nw63[(_dafny.ONE).toNumber()] = (_this).f17;
      _nw63[(new BigNumber(2)).toNumber()] = (_this).f17;
      _nw63[(new BigNumber(3)).toNumber()] = (_this).f17;
      _525_v52 = _nw63;
      let _526_v53;
      _526_v53 = new _dafny.CodePoint('o'.codePointAt(0));
      let _527_v54;
      _527_v54 = _module.D1.create_DC4(_525_v52, _444_v0, _526_v53);
      let _528_v55;
      _528_v55 = _dafny.MultiSet.fromElements(_525_v52, _525_v52);
      let _529_v56;
      _529_v56 = _dafny.Seq.of(_528_v55, _528_v55, _528_v55);
      r0 = ((_527_v54).dtor_cf6).minus(new BigNumber(((_529_v56)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_529_v56).length))]).cardinality()));
      let _530_v57;
      let _init7 = function (_531_i6) {
        return _module.D0.create_DC0(true);
      };
      let _nw64 = Array((new BigNumber(19)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw64.length); _i0_7++) {
        _nw64[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _530_v57 = _nw64;
      let _532_v58;
      _532_v58 = _module.D0.create_DC0((_this).f17);
      let _index99 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_530_v57).length));
      (_530_v57)[_index99] = _532_v58;
      let _index100 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_530_v57).length));
      (_530_v57)[_index100] = _532_v58;
      r2 = !_dafny.Seq.contains(_this.f18, _526_v53);
      let _index101 = _module.__default.safeIndex(new BigNumber(292), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index101] = _444_v0;
      let _index102 = _module.__default.safeIndex(new BigNumber(292), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index102] = new BigNumber(-219);
      let _533_v59;
      _533_v59 = _dafny.Set.fromElements((_this).f17);
      let _534_v60;
      _534_v60 = _dafny.MultiSet.fromElements(p0, (_533_v59).IsDisjointFrom(_533_v59), (true) || (_module.__default.fm0(_447_v4, ((_this).f14)[_module.__default.safeIndex(new BigNumber(292), new BigNumber(((_this).f14).length))], globalState)));
      r0 = (((_534_v60).contains(true)) ? ((_534_v60).get(true)) : (new BigNumber(870)));
      r1 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(292), new BigNumber(((_this).f14).length))];
      r2 = p0;
      return [r0, r1, r2];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    set f13(value) {
      let _this = this;
      _this._f13 = value;
    };
    __ctor(f13) {
      let _this = this;
      (_this)._f13 = f13;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f13;
    };
    fm5(globalState) {
      let _this = this;
      let _source11 = _module.D5.create_DC13();
      if (_source11.is_DC13) {
        return _module.D1.create_DC3(new BigNumber(614));
      } else if (_source11.is_DC14) {
        let _535___mcc_h0 = (_source11).cf23;
        let _536_cf23 = _535___mcc_h0;
        return _module.D1.create_DC3(new BigNumber(504));
      } else if (_source11.is_DC12) {
        let _537___mcc_h1 = (_source11).cf22;
        let _538_cf22 = _537___mcc_h1;
        return _module.D1.create_DC3((((_538_cf22).contains(_this.f13)) ? ((_538_cf22).get(_this.f13)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13)).length))));
      } else {
        let _539___mcc_h2 = (_source11).cf24;
        let _540_cf24 = _539___mcc_h2;
        return _module.D1.create_DC3(new BigNumber((_dafny.Set.fromElements(_this.f13)).length));
      }
    };
    fm29(p0, p1, p2, globalState) {
      let _this = this;
      if (!(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f13,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(248),_this.f13)).length))).length),new BigNumber((_dafny.Seq.UnicodeFromString("eaogdxqm")).length))).equals(function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_dafny.Seq.of(new BigNumber(214), new BigNumber((_dafny.Seq.UnicodeFromString("bqw")).length))).Elements) {
          let _541_v0 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(214), new BigNumber((_dafny.Seq.UnicodeFromString("bqw")).length)), _541_v0)) {
            _coll28.push([(_541_v0).plus(new BigNumber(24)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_this.f13)).length),new BigNumber(-428))).length))]);
          }
        }
        return _coll28;
      }())) {
        return _module.D5.create_DC13();
      } else {
        return _module.D5.create_DC13();
      }
    };
    fm30(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f13;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let _542_v0;
      _542_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
      let _hi3 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
      for (let _543_i0 = (new BigNumber((_542_v0).length)).plus(p0); _543_i0.isLessThan(_hi3); _543_i0 = _543_i0.plus(_dafny.ONE)) {
        if (true) {
          let _544_v1;
          _544_v1 = new BigNumber(661);
          let _545_v2;
          _545_v2 = new _dafny.CodePoint('v'.codePointAt(0));
          let _546_v3;
          _546_v3 = _dafny.Map.Empty.slice().updateUnsafe(_545_v2,_543_i0);
          let _547_v4;
          _547_v4 = _dafny.Seq.of(new BigNumber((_546_v3).length));
          _544_v1 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_547_v4).length)), (_module.D2.create_DC7(_dafny.Seq.of(_545_v2, _545_v2), p0, _544_v1, false)).dtor_cf14);
          let _548_v5;
          _548_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _549_v6;
          _549_v6 = _dafny.Seq.of(_this.f13);
          let _550_v7;
          _550_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm31(globalState),(_549_v6)[_module.__default.safeIndex(_543_i0, new BigNumber((_549_v6).length))]);
          let _551_v8;
          _551_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_548_v5).contains(_543_i0)) ? ((_548_v5).get(_543_i0)) : (_module.__default.fm31(globalState))),_550_v7);
          _551_v8 = (_551_v8).update(_543_i0, _550_v7);
          let _552_v9;
          _552_v9 = _dafny.Seq.UnicodeFromString("ioqxh");
          let _553_v10;
          _553_v10 = _dafny.Map.Empty.slice().updateUnsafe(_552_v9,!(_this.f13));
          let _554_v11;
          let _out28;
          _out28 = _module.__default.m0(_553_v10, globalState);
          _554_v11 = _out28;
          let _555_v12;
          let _nw65 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _555_v12 = _nw65;
          let _index103 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_555_v12).length));
          (_555_v12)[_index103] = _554_v11;
          let _index104 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_555_v12).length));
          (_555_v12)[_index104] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_543_i0));
          (_this).f13 = _this.f13;
        } else {
          (_this).f13 = _this.f13;
          let _556_v13;
          _556_v13 = _dafny.Seq.UnicodeFromString("fnbb");
          let _557_v14;
          let _nw66 = Array((new BigNumber(17)).toNumber()).fill(false);
          _557_v14 = _nw66;
          let _558_v15;
          _558_v15 = new _dafny.CodePoint('j'.codePointAt(0));
          let _559_v16;
          _559_v16 = _module.D1.create_DC4(_557_v14, p0, _558_v15);
          let _560_v17;
          _560_v17 = _dafny.Set.fromElements(_this.f13);
          let _561_v18;
          _561_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_560_v17);
          let _562_v19;
          _562_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_560_v17);
          let _563_v20;
          _563_v20 = _module.D1.create_DC5(_556_v13, (_559_v16).dtor_cf5, new BigNumber(((((_561_v18).contains(p0)) ? ((_561_v18).get(p0)) : ((((_562_v19).contains(_this.f13)) ? ((_562_v19).get(_this.f13)) : (_560_v17))))).length));
          let _564_v21;
          _564_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new BigNumber((_542_v0).length));
          let _565_v22;
          _565_v22 = _dafny.Seq.of(_556_v13);
          let _pat_let_tv15 = _556_v13;
          let _566_v23;
          let _nw67 = Array((new BigNumber(23)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _563_v20;
          _nw67[(_dafny.ONE).toNumber()] = _563_v20;
          _nw67[(new BigNumber(2)).toNumber()] = function (_pat_let12_0) {
            return function (_567_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_568_dt__update_hcf8_h0) {
                  return function (_pat_let14_0) {
                    return function (_569_dt__update_hcf10_h0) {
                      return _module.D1.create_DC5(_568_dt__update_hcf8_h0, (_567_dt__update__tmp_h0).dtor_cf9, _569_dt__update_hcf10_h0);
                    }(_pat_let14_0);
                  }(_543_i0);
                }(_pat_let13_0);
              }(_pat_let_tv15);
            }(_pat_let12_0);
          }(_563_v20);
          _nw67[(new BigNumber(3)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(4)).toNumber()] = _module.D1.create_DC5(_556_v13, _557_v14, _543_i0);
          _nw67[(new BigNumber(5)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(6)).toNumber()] = _module.D1.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), ((_570_v15) => function (_571_i1) {
  return _570_v15;
})(_558_v15)), _557_v14, _543_i0);
          _nw67[(new BigNumber(7)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(8)).toNumber()] = _module.D1.create_DC5(_556_v13, _557_v14, _543_i0);
          _nw67[(new BigNumber(9)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(10)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(11)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(12)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(13)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(14)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(15)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(16)).toNumber()] = _module.D1.create_DC5(_module.__default.fm32(_dafny.Seq.UnicodeFromString("xehdabxv"), _this.f13, (((_564_v21).contains(_this.f13)) ? ((_564_v21).get(_this.f13)) : (p0)), _543_i0, globalState), _557_v14, _543_i0);
          _nw67[(new BigNumber(17)).toNumber()] = _module.D1.create_DC5(_556_v13, _557_v14, new BigNumber((_565_v22).length));
          _nw67[(new BigNumber(18)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(19)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(20)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(21)).toNumber()] = _563_v20;
          _nw67[(new BigNumber(22)).toNumber()] = _563_v20;
          _566_v23 = _nw67;
          let _rhs77 = ((_this.f13) ? (_566_v23) : (_566_v23));
          let _rhs78 = _this.f13;
          let _lhs48 = _this;
          _566_v23 = _rhs77;
          _lhs48.f13 = _rhs78;
          let _572_v24;
          let _init8 = ((_573_i0) => function (_574_i2) {
            return (_574_i2).plus((_dafny.ZERO).minus(_573_i0));
          })(_543_i0);
          let _nw68 = Array((new BigNumber(27)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw68.length); _i0_8++) {
            _nw68[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _572_v24 = _nw68;
          let _index105 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_572_v24).length));
          (_572_v24)[_index105] = _module.__default.fm31(globalState);
          let _index106 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_572_v24).length));
          (_572_v24)[_index106] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), ((_575_v15) => function (_576_i3) {
            return _575_v15;
          })(_558_v15)), _556_v13), _dafny.Seq.Concat(_556_v13, _556_v13))).length);
          let _index107 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_572_v24).length));
          (_572_v24)[_index107] = (_572_v24)[_module.__default.safeIndex(new BigNumber(322), new BigNumber((_572_v24).length))];
          (_this).f13 = _this.f13;
        }
        let _577_v25;
        _577_v25 = _dafny.Seq.UnicodeFromString("hujnpvv");
        let _578_v26;
        _578_v26 = new _dafny.CodePoint('w'.codePointAt(0));
        (_this).f13 = !(_dafny.Seq.IsPrefixOf(_577_v25, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_577_v25, _577_v25), _module.__default.safeIndex(_543_i0, new BigNumber((_dafny.Seq.Concat(_577_v25, _577_v25)).length)), _578_v26), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_577_v25, _577_v25), _module.__default.safeIndex(_543_i0, new BigNumber((_dafny.Seq.Concat(_577_v25, _577_v25)).length)), _578_v26)).length)), new _dafny.CodePoint('f'.codePointAt(0)))));
        let _579_v27;
        _579_v27 = new BigNumber(918);
        _579_v27 = _543_i0;
        _579_v27 = p0;
      }
      let _580_v28;
      let _nw69 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _580_v28 = _nw69;
      let _581_v29;
      _581_v29 = _dafny.Map.Empty.slice().updateUnsafe(_580_v28,_580_v28);
      _581_v29 = (_dafny.Map.Empty.slice().updateUnsafe(_580_v28,_580_v28)).Merge((_581_v29).Merge(_581_v29));
      let _582_v30;
      let _nw70 = new _module.C0();
      _nw70.__ctor(p0);
      _582_v30 = _nw70;
      let _583_v31;
      let _init9 = function (_584_i4) {
        return _this.f13;
      };
      let _nw71 = Array((new BigNumber(29)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw71.length); _i0_9++) {
        _nw71[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _583_v31 = _nw71;
      let _585_v32;
      _585_v32 = _module.D1.create_DC4(_583_v31, (_582_v30).f15, new _dafny.CodePoint('m'.codePointAt(0)));
      let _586_v33;
      _586_v33 = _dafny.Seq.UnicodeFromString("nqogre");
      let _587_v34;
      _587_v34 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (new _dafny.CodePoint('t'.codePointAt(0))) : ((_585_v32).dtor_cf7)),_586_v33);
      let _588_v35;
      _588_v35 = new _dafny.CodePoint('f'.codePointAt(0));
      _587_v34 = (_587_v34).update(_588_v35, _586_v33);
      let _589_v36;
      _589_v36 = _dafny.Seq.of(_this.f13);
      let _index108 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_580_v28).length));
      (_580_v28)[_index108] = new BigNumber((_589_v36).length);
      let _index109 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_580_v28).length));
      (_580_v28)[_index109] = (_582_v30).f15;
      let _590_v37;
      _590_v37 = _dafny.MultiSet.fromElements((_580_v28)[_module.__default.safeIndex(new BigNumber(501), new BigNumber((_580_v28).length))]);
      _589_v36 = _dafny.Seq.of((_590_v37).IsProperSubsetOf(_590_v37), !((new BigNumber(523)).isLessThan(p0)));
      r0 = _580_v28;
      return r0;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _591_v0;
      _591_v0 = _dafny.Set.fromElements(p0);
      let _hi4 = p2;
      for (let _592_i0 = new BigNumber((_module.__default.fm33(_this.f13, false, _591_v0, _this.f13, globalState)).length); _592_i0.isLessThan(_hi4); _592_i0 = _592_i0.plus(_dafny.ONE)) {
        let _593_v1;
        let _nw72 = Array((new BigNumber(24)).toNumber()).fill([]);
        _593_v1 = _nw72;
        let _594_v2;
        let _nw73 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _594_v2 = _nw73;
        let _595_v3;
        _595_v3 = new _dafny.CodePoint('p'.codePointAt(0));
        let _596_v4;
        _596_v4 = _dafny.MultiSet.fromElements(_595_v3);
        let _index110 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_594_v2).length));
        (_594_v2)[_index110] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("wawnlrb")).length), (((_596_v4).contains(_595_v3)) ? ((_596_v4).get(_595_v3)) : (_592_i0)));
        let _597_v5;
        _597_v5 = _dafny.Set.fromElements(_this.f13, _this.f13, _this.f13);
        let _598_v6;
        _598_v6 = _dafny.Seq.of(p0, _module.__default.fm31(globalState), new BigNumber((_597_v5).length), p0);
        let _599_v7;
        _599_v7 = _dafny.MultiSet.fromElements(_598_v6, _598_v6);
        let _600_v8;
        _600_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_592_i0);
        let _601_v9;
        _601_v9 = _module.D5.create_DC12(_600_v8);
        let _602_v10;
        _602_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_601_v9);
        let _603_v11;
        _603_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm30(_599_v7, p0, _602_v10, globalState),new _dafny.CodePoint('h'.codePointAt(0)));
        let _index111 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_594_v2).length));
        let _rhs79 = _this.f13;
        let _rhs80 = _593_v1;
        let _rhs81 = (new BigNumber((_dafny.Seq.UnicodeFromString("tasrqgp")).length)).minus(new BigNumber(((_603_v11).Merge(_603_v11)).length));
        let _lhs49 = _594_v2;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_594_v2).length));
        r2 = _rhs79;
        _593_v1 = _rhs80;
        _lhs49[_lhs50] = _rhs81;
        r1 = (((_594_v2)[_module.__default.safeIndex(new BigNumber(110), new BigNumber((_594_v2).length))]).isLessThanOrEqualTo(_592_i0)) || (_this.f13);
        let _index112 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_594_v2).length));
        (_594_v2)[_index112] = new BigNumber(62);
        let _604_v12;
        let _nw74 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
        _604_v12 = _nw74;
        let _605_v13;
        let _nw75 = Array((new BigNumber(4)).toNumber()).fill(false);
        _605_v13 = _nw75;
        let _606_v14;
        _606_v14 = _dafny.MultiSet.fromElements(_605_v13);
        let _index113 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_604_v12).length));
        (_604_v12)[_index113] = _606_v14;
        let _index114 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_604_v12).length));
        (_604_v12)[_index114] = (_606_v14).update(_605_v13, _module.__default.abs(_592_i0));
      }
      let _607_v16;
      _607_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,p0);
      let _608_v17;
      _608_v17 = new _dafny.CodePoint('w'.codePointAt(0));
      let _609_v18;
      _609_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_607_v16).length), new BigNumber((p1).length)), _608_v17),p2);
      let _610_v19;
      let _out29;
      _out29 = _module.__default.m0(function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_609_v18).Keys.Elements) {
          let _611_v15 = _compr_29;
          if ((_609_v18).contains(_611_v15)) {
            _coll29.push([_611_v15,_this.f13]);
          }
        }
        return _coll29;
      }(), globalState);
      _610_v19 = _out29;
      let _612_i1;
      _612_i1 = _dafny.ZERO;
      L7: {
        while (_this.f13) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_612_i1)) {
              break L7;
            }
            _612_i1 = (_612_i1).plus(_dafny.ONE);
            let _613_v20;
            _613_v20 = _dafny.MultiSet.fromElements(new BigNumber(251));
            let _614_v21;
            _614_v21 = _module.D2.create_DC7(p1, p2, _610_v19, !(_this.f13));
            let _615_v22;
            _615_v22 = _dafny.Set.fromElements(_608_v17);
            let _616_v23;
            _616_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_615_v22).length),p2);
            let _617_v24;
            _617_v24 = _dafny.Map.Empty.slice().updateUnsafe(_613_v20,(_dafny.ZERO).minus(_610_v19));
            let _618_v25;
            _618_v25 = _module.D0.create_DC1(_this.f13, p2);
            let _619_v26;
            let _nw76 = Array((new BigNumber(5)).toNumber());
            _nw76[(_dafny.ZERO).toNumber()] = _this.f13;
            _nw76[(_dafny.ONE).toNumber()] = !_dafny.Seq.contains(p1, _module.__default.fm34(globalState));
            _nw76[(new BigNumber(2)).toNumber()] = !(_617_v24).contains((_613_v20).update(new BigNumber(((_614_v21).dtor_cf12).length), _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_616_v23).length)))));
            _nw76[(new BigNumber(3)).toNumber()] = (_this).fm4(_this.f13, _618_v25, _610_v19, _616_v23, globalState);
            _nw76[(new BigNumber(4)).toNumber()] = _this.f13;
            _619_v26 = _nw76;
            let _index115 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length));
            (_619_v26)[_index115] = true;
            let _index116 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length));
            (_619_v26)[_index116] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_620_v17) => function (_621_i2) {
              return _620_v17;
            })(_608_v17)), _dafny.Seq.UnicodeFromString("k"));
            let _622_v27;
            _622_v27 = _dafny.Seq.UnicodeFromString("y");
            _622_v27 = _622_v27;
            if ((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]) {
              let _623_v28;
              let _nw77 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
              _623_v28 = _nw77;
              let _nw78 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
              _623_v28 = _nw78;
              _610_v19 = (_module.__default.fm31(globalState)).multipliedBy((new BigNumber((_622_v27).length)).multipliedBy(_module.__default.fm31(globalState)));
              let _624_v29;
              _624_v29 = _dafny.Seq.of(!((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]), (_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]);
              let _625_v30;
              _625_v30 = _dafny.Set.fromElements((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]);
              let _626_v31;
              _626_v31 = _dafny.Seq.of(new BigNumber((_625_v30).length), _610_v19);
              let _627_v32;
              _627_v32 = _module.D1.create_DC4(_619_v26, p2, new _dafny.CodePoint('y'.codePointAt(0)));
              let _628_v33;
              _628_v33 = _dafny.MultiSet.fromElements((_627_v32).dtor_cf7);
              let _629_v34;
              let _nw79 = Array((new BigNumber(15)).toNumber());
              _nw79[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_624_v29)).cardinality());
              _nw79[(_dafny.ONE).toNumber()] = _610_v19;
              _nw79[(new BigNumber(2)).toNumber()] = new BigNumber(300);
              _nw79[(new BigNumber(3)).toNumber()] = p0;
              _nw79[(new BigNumber(4)).toNumber()] = p0;
              _nw79[(new BigNumber(5)).toNumber()] = new BigNumber((_624_v29).length);
              _nw79[(new BigNumber(6)).toNumber()] = new BigNumber((_607_v16).length);
              _nw79[(new BigNumber(7)).toNumber()] = p2;
              _nw79[(new BigNumber(8)).toNumber()] = p2;
              _nw79[(new BigNumber(9)).toNumber()] = (_626_v31)[_module.__default.safeIndex((((_628_v33).contains(_608_v17)) ? ((_628_v33).get(_608_v17)) : (_610_v19)), new BigNumber((_626_v31).length))];
              _nw79[(new BigNumber(10)).toNumber()] = _610_v19;
              _nw79[(new BigNumber(11)).toNumber()] = p0;
              _nw79[(new BigNumber(12)).toNumber()] = (((_613_v20).contains(_610_v19)) ? ((_613_v20).get(_610_v19)) : (p0));
              _nw79[(new BigNumber(13)).toNumber()] = new BigNumber(304);
              _nw79[(new BigNumber(14)).toNumber()] = p2;
              _629_v34 = _nw79;
              let _630_v35;
              let _nw80 = new _module.C1();
              _nw80.__ctor(_this.f13, _622_v27, _629_v34);
              _630_v35 = _nw80;
              _610_v19 = p0;
              let _631_v36;
              _631_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_608_v17);
              let _632_v37;
              let _nw81 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
              _632_v37 = _nw81;
              let _index117 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_632_v37).length));
              (_632_v37)[_index117] = _624_v29;
              let _index118 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_632_v37).length));
              let _rhs82 = new BigNumber(-25);
              let _rhs83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(p1, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("krd"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("krd")).length)), new _dafny.CodePoint('o'.codePointAt(0)))),_608_v17);
              let _rhs84 = _626_v31;
              let _rhs85 = _dafny.Seq.of((_this.f13) === (_this.f13), (_630_v35).f17);
              let _rhs86 = _622_v27;
              let _lhs51 = _632_v37;
              let _lhs52 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_632_v37).length));
              _610_v19 = _rhs82;
              _631_v36 = _rhs83;
              _626_v31 = _rhs84;
              _lhs51[_lhs52] = _rhs85;
              _622_v27 = _rhs86;
            } else {
              let _633_v38;
              let _nw82 = Array((new BigNumber(3)).toNumber());
              _nw82[(_dafny.ZERO).toNumber()] = p2;
              _nw82[(_dafny.ONE).toNumber()] = _610_v19;
              _nw82[(new BigNumber(2)).toNumber()] = new BigNumber(149);
              _633_v38 = _nw82;
              let _index119 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_633_v38).length));
              (_633_v38)[_index119] = _610_v19;
              let _634_v39;
              _634_v39 = _dafny.Seq.of(_610_v19, p0, p0);
              let _index120 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_633_v38).length));
              (_633_v38)[_index120] = (((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]) ? ((p2).multipliedBy(p2)) : (new BigNumber((_dafny.Seq.Concat(_634_v39, _634_v39)).length)));
              let _635_v40;
              let _nw83 = new _module.C0();
              _nw83.__ctor(_610_v19);
              _635_v40 = _nw83;
              let _636_v41;
              let _nw84 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
              _636_v41 = _nw84;
              _636_v41 = _636_v41;
              let _rhs87 = _608_v17;
              let _rhs88 = _this.f13;
              let _lhs53 = _this;
              _608_v17 = _rhs87;
              _lhs53.f13 = _rhs88;
              _622_v27 = (((((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]) ? ((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]) : ((_619_v26)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_619_v26).length))]))) ? (p1) : (p1));
            }
            let _637_v42;
            _637_v42 = _dafny.Map.Empty.slice().updateUnsafe(_622_v27,_619_v26);
            _637_v42 = (_637_v42).update(_622_v27, _619_v26);
          }
        }
      }
      let _hi5 = new BigNumber(151);
      for (let _638_i3 = (_dafny.ZERO).minus(new BigNumber(((_591_v0).Difference(_591_v0)).length)); _638_i3.isLessThan(_hi5); _638_i3 = _638_i3.plus(_dafny.ONE)) {
        let _639_v43;
        _639_v43 = _dafny.Seq.of(_this.f13);
        let _640_v44;
        _640_v44 = _dafny.Seq.of((_639_v43)[_module.__default.safeIndex(p2, new BigNumber((_639_v43).length))]);
        let _641_v45;
        _641_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),_this.f13);
        (_this).f13 = ((_641_v45).contains(_639_v43)) && (_this.f13);
        let _642_v46;
        let _nw85 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _642_v46 = _nw85;
        let _643_v47;
        let _nw86 = new _module.C1();
        _nw86.__ctor(_this.f13, _dafny.Seq.UnicodeFromString("gwxskgni"), _642_v46);
        _643_v47 = _nw86;
        let _644_v48;
        _644_v48 = _module.D0.create_DC0(_this.f13);
        let _pat_let_tv16 = _643_v47;
        let _pat_let_tv17 = _643_v47;
        let _645_v49;
        let _nw87 = Array((new BigNumber(6)).toNumber());
        _nw87[(_dafny.ZERO).toNumber()] = function (_pat_let15_0) {
          return function (_646_dt__update__tmp_h0) {
            return function (_pat_let16_0) {
              return function (_647_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_647_dt__update_hcf0_h0);
              }(_pat_let16_0);
            }(false);
          }(_pat_let15_0);
        }(_644_v48);
        _nw87[(_dafny.ONE).toNumber()] = function (_pat_let17_0) {
          return function (_648_dt__update__tmp_h1) {
            return function (_pat_let18_0) {
              return function (_649_dt__update_hcf0_h1) {
                return _module.D0.create_DC0(_649_dt__update_hcf0_h1);
              }(_pat_let18_0);
            }((_pat_let_tv16).f17);
          }(_pat_let17_0);
        }(_644_v48);
        _nw87[(new BigNumber(2)).toNumber()] = _644_v48;
        _nw87[(new BigNumber(3)).toNumber()] = function (_pat_let19_0) {
          return function (_650_dt__update__tmp_h2) {
            return function (_pat_let20_0) {
              return function (_651_dt__update_hcf0_h2) {
                return _module.D0.create_DC0(_651_dt__update_hcf0_h2);
              }(_pat_let20_0);
            }((_pat_let_tv17).f17);
          }(_pat_let19_0);
        }(_644_v48);
        _nw87[(new BigNumber(4)).toNumber()] = _644_v48;
        _nw87[(new BigNumber(5)).toNumber()] = _module.D0.create_DC0(_this.f13);
        _645_v49 = _nw87;
        (_643_v47).m5(_645_v49, globalState);
        let _652_v50;
        _652_v50 = _dafny.Map.Empty.slice().updateUnsafe((_643_v47).f17,_dafny.Map.Empty.slice().updateUnsafe(_638_i3,_638_i3));
        let _653_v51;
        _653_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_639_v43),!(_652_v50).contains((_643_v47).f17));
        let _654_v52;
        _654_v52 = _dafny.MultiSet.fromElements(true);
        let _655_v53;
        let _nw88 = Array((new BigNumber(20)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = (_643_v47).f17;
        _nw88[(_dafny.ONE).toNumber()] = _this.f13;
        _nw88[(new BigNumber(2)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(3)).toNumber()] = (_643_v47).f17;
        _nw88[(new BigNumber(4)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(5)).toNumber()] = !(_this.f13);
        _nw88[(new BigNumber(6)).toNumber()] = !(_this.f13);
        _nw88[(new BigNumber(7)).toNumber()] = !((_654_v52).IsDisjointFrom(_654_v52));
        _nw88[(new BigNumber(8)).toNumber()] = !((_643_v47).f17);
        _nw88[(new BigNumber(9)).toNumber()] = (_643_v47).f17;
        _nw88[(new BigNumber(10)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(11)).toNumber()] = (_643_v47).f17;
        _nw88[(new BigNumber(12)).toNumber()] = (p0).isEqualTo(p0);
        _nw88[(new BigNumber(13)).toNumber()] = true;
        _nw88[(new BigNumber(14)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(15)).toNumber()] = !(_this.f13) || ((_643_v47).f17);
        _nw88[(new BigNumber(16)).toNumber()] = (p2).isLessThan(p2);
        _nw88[(new BigNumber(17)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(18)).toNumber()] = _this.f13;
        _nw88[(new BigNumber(19)).toNumber()] = !(_this.f13);
        _655_v53 = _nw88;
        let _index121 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_655_v53).length));
        (_655_v53)[_index121] = (_643_v47).fm6(globalState);
        let _index122 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_655_v53).length));
        (_655_v53)[_index122] = _this.f13;
        let _656_v54;
        _656_v54 = _dafny.Seq.of(_638_i3, p2);
        let _657_v55;
        _657_v55 = _dafny.Map.Empty.slice().updateUnsafe((_643_v47).fm7(_638_i3, (_dafny.ZERO).minus((_656_v54)[_module.__default.safeIndex(new BigNumber((_643_v47.f18).length), new BigNumber((_656_v54).length))]), _this.f13, globalState),new BigNumber((p1).length));
        let _658_v56;
        _658_v56 = _dafny.Set.fromElements((_643_v47).fm14(_638_i3, (_643_v47).f17, new BigNumber((_657_v55).length), p2, globalState));
        let _index123 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_655_v53).length));
        let _index124 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_655_v53).length));
        let _rhs89 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f13, (_644_v48).dtor_cf0, (_643_v47).f17),(_643_v47).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_640_v44),(_643_v47).f17));
        let _rhs90 = true;
        let _rhs91 = (_658_v56).IsSubsetOf(_658_v56);
        let _rhs92 = !((_643_v47).f17);
        let _rhs93 = !_dafny.areEqual(p1, _643_v47.f18);
        let _lhs54 = _655_v53;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_655_v53).length));
        let _lhs56 = _655_v53;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_655_v53).length));
        _653_v51 = _rhs89;
        r1 = _rhs90;
        _lhs54[_lhs55] = _rhs91;
        _lhs56[_lhs57] = _rhs92;
        r1 = _rhs93;
      }
      let _659_v57;
      _659_v57 = _dafny.Set.fromElements(_this.f13, !(_this.f13));
      let _660_v58;
      _660_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
      let _661_v59;
      let _nw89 = Array((new BigNumber(16)).toNumber());
      _nw89[(_dafny.ZERO).toNumber()] = p0;
      _nw89[(_dafny.ONE).toNumber()] = p2;
      _nw89[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw89[(new BigNumber(3)).toNumber()] = p0;
      _nw89[(new BigNumber(4)).toNumber()] = p0;
      _nw89[(new BigNumber(5)).toNumber()] = new BigNumber((_659_v57).length);
      _nw89[(new BigNumber(6)).toNumber()] = p2;
      _nw89[(new BigNumber(7)).toNumber()] = p0;
      _nw89[(new BigNumber(8)).toNumber()] = p2;
      _nw89[(new BigNumber(9)).toNumber()] = new BigNumber((p1).length);
      _nw89[(new BigNumber(10)).toNumber()] = _610_v19;
      _nw89[(new BigNumber(11)).toNumber()] = p0;
      _nw89[(new BigNumber(12)).toNumber()] = _610_v19;
      _nw89[(new BigNumber(13)).toNumber()] = new BigNumber((_660_v58).length);
      _nw89[(new BigNumber(14)).toNumber()] = _610_v19;
      _nw89[(new BigNumber(15)).toNumber()] = p2;
      _661_v59 = _nw89;
      let _662_v60;
      _662_v60 = _module.D4.create_DC10(((_this.f13) ? (_661_v59) : (_661_v59)));
      let _source12 = _662_v60;
      if (_source12.is_DC11) {
        let _663___mcc_h0 = (_source12).cf21;
        let _664_cf21 = _663___mcc_h0;
        let _665_v61;
        let _nw90 = new _module.C0();
        _nw90.__ctor(_module.__default.safeDivisionInt(new BigNumber(822), _module.__default.fm31(globalState)));
        _665_v61 = _nw90;
        let _666_v62;
        _666_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm35(p0, globalState),false);
        let _667_v63;
        let _nw91 = Array((new BigNumber(14)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _661_v59;
        _nw91[(_dafny.ONE).toNumber()] = _661_v59;
        _nw91[(new BigNumber(2)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(3)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(4)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(5)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(6)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(7)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(8)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(9)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(10)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(11)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(12)).toNumber()] = _661_v59;
        _nw91[(new BigNumber(13)).toNumber()] = _661_v59;
        _667_v63 = _nw91;
        let _668_v64;
        let _nw92 = new _module.C1();
        _nw92.__ctor(_this.f13, p1, _661_v59);
        _668_v64 = _nw92;
        let _669_v65;
        _669_v65 = _dafny.Map.Empty.slice().updateUnsafe(_667_v63,_dafny.Set.fromElements(_668_v64));
        let _670_v66;
        _670_v66 = _dafny.Set.fromElements(_668_v64);
        let _671_v67;
        _671_v67 = _dafny.Map.Empty.slice().updateUnsafe((p2).minus(new BigNumber((_666_v62).length)),((!(_664_cf21)) ? ((_dafny.ZERO).minus(new BigNumber(((((_669_v65).contains(_667_v63)) ? ((_669_v65).get(_667_v63)) : (_670_v66))).length))) : (_610_v19)));
        _671_v67 = _671_v67;
        let _672_v68;
        let _nw93 = Array((new BigNumber(11)).toNumber()).fill(false);
        _672_v68 = _nw93;
        let _index125 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_672_v68).length));
        (_672_v68)[_index125] = ((_this.f13) ? (false) : (_664_cf21));
        let _673_v69;
        _673_v69 = _dafny.MultiSet.fromElements(_672_v68, _672_v68, _672_v68);
        let _index126 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_672_v68).length));
        (_672_v68)[_index126] = !(((_673_v69).IsDisjointFrom(_673_v69)) || (_664_cf21));
        let _index127 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_672_v68).length));
        (_672_v68)[_index127] = (_668_v64).f17;
      } else {
        let _674___mcc_h1 = (_source12).cf20;
        let _675_cf20 = _674___mcc_h1;
        let _676_v70;
        _676_v70 = _dafny.Seq.of(_this.f13);
        let _677_v71;
        let _nw94 = new _module.C1();
        _nw94.__ctor((!((_676_v70)[_module.__default.safeIndex(_610_v19, new BigNumber((_676_v70).length))])) === (_this.f13), p1, _675_cf20);
        _677_v71 = _nw94;
        if (_this.f13) {
          let _678_v72;
          _678_v72 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("cqgr"), _module.__default.safeIndex(_610_v19, new BigNumber((_dafny.Seq.UnicodeFromString("cqgr")).length)), _608_v17), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-754)), ((_679_v17) => function (_680_i4) {
            return _679_v17;
          })(_608_v17)), _dafny.Seq.UnicodeFromString("uthxbo"));
          let _681_v73;
          _681_v73 = _dafny.Map.Empty.slice().updateUnsafe((_677_v71).f17,_678_v72);
          let _682_v74;
          _682_v74 = _dafny.Seq.of(new BigNumber((_659_v57).length));
          r1 = _dafny.Seq.IsPrefixOf(((_this.f13) ? ((((_681_v73).contains((_677_v71).f17)) ? ((_681_v73).get((_677_v71).f17)) : (_678_v72))) : (_678_v72)), _dafny.Seq.of(_module.__default.fm32(_677_v71.f18, (_677_v71).f17, p2, (_682_v74)[_module.__default.safeIndex(new BigNumber((_607_v16).length), new BigNumber((_682_v74).length))], globalState)));
          _610_v19 = (_dafny.ZERO).minus((_610_v19).multipliedBy(_610_v19));
          _591_v0 = (((_677_v71).f17) ? (_591_v0) : (_591_v0));
          _610_v19 = (new BigNumber((_dafny.Seq.of(p2)).length)).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_609_v18).length), new BigNumber(416))));
          let _683_v75;
          let _nw95 = new _module.C1();
          _nw95.__ctor((_677_v71).f17, p1, _675_cf20);
          _683_v75 = _nw95;
        } else {
          let _684_v76;
          let _init10 = ((_685_v17) => function (_686_i5) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), ((_687_v17) => function (_688_i6) {
              return _687_v17;
            })(_685_v17));
          })(_608_v17);
          let _nw96 = Array((new BigNumber(13)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw96.length); _i0_10++) {
            _nw96[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _684_v76 = _nw96;
          let _689_v77;
          _689_v77 = _dafny.Seq.of(_684_v76, _684_v76, _684_v76, _684_v76, _684_v76);
          _684_v76 = (_689_v77)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_610_v19,true)).length), new BigNumber((_689_v77).length))];
          let _index128 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_675_cf20).length));
          (_675_cf20)[_index128] = p0;
          let _index129 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_675_cf20).length));
          (_675_cf20)[_index129] = _module.__default.safeModuloInt(p2, p0);
          r2 = _this.f13;
          _610_v19 = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(-336)).minus((_675_cf20)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_675_cf20).length))])));
          let _690_v78;
          _690_v78 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f13);
          _690_v78 = (_690_v78).Merge(_690_v78);
        }
        r2 = (_677_v71).f17;
        let _source13 = (((_610_v19).isLessThan(p2)) ? (_662_v60) : (_module.D4.create_DC10(_675_cf20)));
        if (_source13.is_DC11) {
          let _691___mcc_h2 = (_source13).cf21;
          let _692_cf21 = _691___mcc_h2;
          let _693_v79;
          let _nw97 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _693_v79 = _nw97;
          let _694_v80;
          _694_v80 = _dafny.Seq.of(p0);
          let _695_v81;
          _695_v81 = _dafny.Seq.of((_694_v80)[_module.__default.safeIndex(_610_v19, new BigNumber((_694_v80).length))]);
          let _index130 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_693_v79).length));
          (_693_v79)[_index130] = _dafny.Seq.Concat(_695_v81, _694_v80);
          let _696_v82;
          let _nw98 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _696_v82 = _nw98;
          let _697_v83;
          _697_v83 = _module.D0.create_DC1((_677_v71).f17, p0);
          let _698_v84;
          _698_v84 = _module.D5.create_DC14(_692_cf21);
          let _699_v85;
          _699_v85 = _module.D4.create_DC11(_this.f13);
          let _700_v86;
          let _nw99 = Array((new BigNumber(26)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = (_677_v71).f17;
          _nw99[(_dafny.ONE).toNumber()] = _this.f13;
          _nw99[(new BigNumber(2)).toNumber()] = _692_cf21;
          _nw99[(new BigNumber(3)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(4)).toNumber()] = false;
          _nw99[(new BigNumber(5)).toNumber()] = _692_cf21;
          _nw99[(new BigNumber(6)).toNumber()] = false;
          _nw99[(new BigNumber(7)).toNumber()] = !(!((_677_v71).f17));
          _nw99[(new BigNumber(8)).toNumber()] = _692_cf21;
          _nw99[(new BigNumber(9)).toNumber()] = false;
          _nw99[(new BigNumber(10)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(11)).toNumber()] = (_this).fm4((_677_v71).f17, _697_v83, new BigNumber((_module.__default.fm36(globalState)).cardinality()), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(106),_610_v19), globalState);
          _nw99[(new BigNumber(12)).toNumber()] = _692_cf21;
          _nw99[(new BigNumber(13)).toNumber()] = !((_677_v71).fm14(p0, (_677_v71).f17, _610_v19, (_677_v71).fm7(p2, new BigNumber((_677_v71.f18).length), _this.f13, globalState), globalState));
          _nw99[(new BigNumber(14)).toNumber()] = (_698_v84).dtor_cf23;
          _nw99[(new BigNumber(15)).toNumber()] = _this.f13;
          _nw99[(new BigNumber(16)).toNumber()] = _this.f13;
          _nw99[(new BigNumber(17)).toNumber()] = (_676_v70)[_module.__default.safeIndex(_610_v19, new BigNumber((_676_v70).length))];
          _nw99[(new BigNumber(18)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(19)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(20)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(21)).toNumber()] = !((_677_v71).f17);
          _nw99[(new BigNumber(22)).toNumber()] = (_677_v71).fm6(globalState);
          _nw99[(new BigNumber(23)).toNumber()] = (_677_v71).f17;
          _nw99[(new BigNumber(24)).toNumber()] = true;
          _nw99[(new BigNumber(25)).toNumber()] = (_699_v85).dtor_cf21;
          _700_v86 = _nw99;
          let _index131 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_696_v82).length));
          (_696_v82)[_index131] = _dafny.Map.Empty.slice().updateUnsafe(p0,_700_v86);
          let _index132 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_700_v86).length));
          (_700_v86)[_index132] = (_677_v71).f17;
          let _701_v87;
          _701_v87 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2));
          let _702_v88;
          _702_v88 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_701_v87).cardinality())).multipliedBy(p2),_700_v86);
          let _index133 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_693_v79).length));
          let _index134 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_696_v82).length));
          let _index135 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_700_v86).length));
          let _rhs94 = _module.__default.fm35((_677_v71).fm15(globalState), globalState);
          let _rhs95 = _661_v59;
          let _rhs96 = _702_v88;
          let _rhs97 = (_this.f13) === (_this.f13);
          let _rhs98 = _675_cf20;
          let _lhs58 = _693_v79;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_693_v79).length));
          let _lhs60 = _696_v82;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_696_v82).length));
          let _lhs62 = _700_v86;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_700_v86).length));
          _lhs58[_lhs59] = _rhs94;
          _661_v59 = _rhs95;
          _lhs60[_lhs61] = _rhs96;
          _lhs62[_lhs63] = _rhs97;
          _675_cf20 = _rhs98;
          _610_v19 = (((_607_v16).contains(_this.f13)) ? ((_607_v16).get(_this.f13)) : (p0));
          let _703_v89;
          _703_v89 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_675_cf20);
          let _704_v90;
          _704_v90 = _dafny.Map.Empty.slice().updateUnsafe(p2,(p0).isLessThan(new BigNumber((_703_v89).length)));
          let _rhs99 = (_677_v71).fm6(globalState);
          let _rhs100 = _659_v57;
          let _rhs101 = (p2).plus(((false) ? (p0) : ((_677_v71).fm7(_610_v19, p2, (_677_v71).f17, globalState))));
          let _rhs102 = (((_704_v90).contains(_module.__default.safeDivisionInt(p0, _610_v19))) ? ((_704_v90).get(_module.__default.safeDivisionInt(p0, _610_v19))) : ((_701_v87).IsSubsetOf(_701_v87)));
          let _rhs103 = _this.f13;
          r2 = _rhs99;
          _659_v57 = _rhs100;
          _610_v19 = _rhs101;
          _692_cf21 = _rhs102;
          r1 = _rhs103;
          _610_v19 = new BigNumber(976);
        } else {
          let _705___mcc_h3 = (_source13).cf20;
          let _706_cf20 = _705___mcc_h3;
          _610_v19 = new BigNumber((((_660_v58).Merge(_dafny.Map.Empty.slice().updateUnsafe((_677_v71).f17,(_677_v71).f17))).Merge((_660_v58).Merge(_dafny.Map.Empty.slice().updateUnsafe((_677_v71).f17,(_677_v71).f17)))).length);
          _610_v19 = _610_v19;
          let _707_v91;
          _707_v91 = _dafny.MultiSet.fromElements(_this.f13);
          _610_v19 = _module.__default.safeDivisionInt((p2).multipliedBy(new BigNumber(((_707_v91).update((_677_v71).f17, _module.__default.abs(p0))).cardinality())), _610_v19);
          let _708_v92;
          _708_v92 = _dafny.Seq.of(_677_v71);
          let _709_v93;
          _709_v93 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(581));
          let _710_v94;
          _710_v94 = _dafny.Map.Empty.slice().updateUnsafe((((_709_v93).contains(p0)) ? ((_709_v93).get(p0)) : ((((_709_v93).contains(new BigNumber(26))) ? ((_709_v93).get(new BigNumber(26))) : (p2)))),_dafny.Seq.UnicodeFromString("t"));
          let _711_v95;
          let _nw100 = new _module.C1();
          _nw100.__ctor(((_677_v71).fm15(globalState)).isLessThanOrEqualTo(new BigNumber((_708_v92).length)), (((_710_v94).contains(_610_v19)) ? ((_710_v94).get(_610_v19)) : (p1)), _706_cf20);
          _711_v95 = _nw100;
          let _712_v96;
          _712_v96 = _dafny.Map.Empty.slice().updateUnsafe(_610_v19,_607_v16);
          let _index136 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_711_v95).f14).length));
          ((_711_v95).f14)[_index136] = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_677_v71.f18, _677_v71.f18, _dafny.Seq.UnicodeFromString("sp"), _677_v71.f18, _677_v71.f18), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_677_v71.f18, _677_v71.f18, _dafny.Seq.UnicodeFromString("sp"), _677_v71.f18, _677_v71.f18)).length)), p1), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_677_v71.f18, _677_v71.f18, _dafny.Seq.UnicodeFromString("sp"), _677_v71.f18, _677_v71.f18), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_677_v71.f18, _677_v71.f18, _dafny.Seq.UnicodeFromString("sp"), _677_v71.f18, _677_v71.f18)).length)), p1)).length)), _677_v71.f18)).length);
          let _713_v97;
          _713_v97 = _module.D7.create_DC19(_711_v95);
          let _index137 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_711_v95).f14).length));
          let _rhs104 = (_713_v97).dtor_cf30;
          let _rhs105 = _module.__default.fm37((_677_v71).f17, globalState);
          let _rhs106 = _module.__default.safeDivisionInt(p2, (new BigNumber((p1).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-849)), ((_714_v17) => function (_715_i7) {
            return _714_v17;
          })(_608_v17))).length)));
          let _lhs64 = (_711_v95).f14;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_711_v95).f14).length));
          _711_v95 = _rhs104;
          _712_v96 = _rhs105;
          _lhs64[_lhs65] = _rhs106;
        }
      }
      let _716_v98;
      let _init11 = function (_717_i8) {
        return _this.f13;
      };
      let _nw101 = Array((new BigNumber(24)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw101.length); _i0_11++) {
        _nw101[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _716_v98 = _nw101;
      let _718_v99;
      _718_v99 = _dafny.MultiSet.fromElements(_716_v98, _716_v98);
      (_this).f13 = !((_718_v99).Union(_718_v99)).equals((_dafny.MultiSet.fromElements(_716_v98)).Intersect(_718_v99));
      let _719_v100;
      _719_v100 = _dafny.Map.Empty.slice().updateUnsafe(p2,_716_v98);
      r0 = (_719_v100).Merge(_719_v100);
      r1 = _this.f13;
      let _720_v101;
      _720_v101 = _module.D2.create_DC7(_dafny.Seq.of(_608_v17), new BigNumber((_659_v57).length), new BigNumber(749), _this.f13);
      r2 = (((_this.f13) ? (_720_v101) : (_720_v101))).dtor_cf15;
      let _721_v102;
      _721_v102 = _dafny.Seq.of(_this.f13);
      r3 = ((_dafny.Seq.contains(_721_v102, _this.f13)) ? (_721_v102) : (_dafny.Seq.Concat(_721_v102, _721_v102)));
      return [r0, r1, r2, r3];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f13 = false;
      this._f19 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    set f13(value) {
      let _this = this;
      _this._f13 = value;
    };
    __ctor(f19, f13) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f13 = f13;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source14 = _module.D1.create_DC2(function () {
  let _coll30 = new _dafny.Map();
  for (const _compr_30 of _dafny.IntegerRange(new BigNumber(439), new BigNumber(-49))) {
    let _722_v0 = _compr_30;
    if (((new BigNumber(439)).isLessThanOrEqualTo(_722_v0)) && ((_722_v0).isLessThan(new BigNumber(-49)))) {
      _coll30.push([(_722_v0).plus(new BigNumber(-610)),new BigNumber(836)]);
    }
  }
  return _coll30;
}());
      if (_source14.is_DC3) {
        let _723___mcc_h0 = (_source14).cf4;
        let _724_cf4 = _723___mcc_h0;
        return true;
      } else if (_source14.is_DC4) {
        let _725___mcc_h1 = (_source14).cf5;
        let _726___mcc_h2 = (_source14).cf6;
        let _727___mcc_h3 = (_source14).cf7;
        let _728_cf7 = _727___mcc_h3;
        let _729_cf6 = _726___mcc_h2;
        let _730_cf5 = _725___mcc_h1;
        return _this.f13;
      } else if (_source14.is_DC5) {
        let _731___mcc_h4 = (_source14).cf8;
        let _732___mcc_h5 = (_source14).cf9;
        let _733___mcc_h6 = (_source14).cf10;
        let _734_cf10 = _733___mcc_h6;
        let _735_cf9 = _732___mcc_h5;
        let _736_cf8 = _731___mcc_h4;
        return false;
      } else {
        let _737___mcc_h7 = (_source14).cf3;
        let _738_cf3 = _737___mcc_h7;
        return _this.f13;
      }
    };
    fm5(globalState) {
      let _this = this;
      if (_this.f13) {
        return _module.D1.create_DC3(new BigNumber((_dafny.Seq.of((_this).f19, _this.f13)).length));
      } else {
        return _module.D1.create_DC3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f13),(_this).f19)).length));
      }
    };
    fm27(globalState) {
      let _this = this;
      return new _dafny.CodePoint('d'.codePointAt(0));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let _hi6 = p0;
      for (let _739_i0 = p0; _739_i0.isLessThan(_hi6); _739_i0 = _739_i0.plus(_dafny.ONE)) {
        let _740_v0;
        _740_v0 = _dafny.Seq.of(_this.f13, (_this).f19);
        (_this).f13 = (_740_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), function (_741_i1) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length), new BigNumber((_740_v0).length))];
        let _742_v1;
        let _nw102 = Array((new BigNumber(28)).toNumber()).fill([]);
        _742_v1 = _nw102;
        let _743_v2;
        let _nw103 = Array((new BigNumber(24)).toNumber()).fill(false);
        _743_v2 = _nw103;
        let _index138 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_742_v1).length));
        (_742_v1)[_index138] = _743_v2;
        let _index139 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_742_v1).length));
        let _init12 = function (_744_i2) {
          return false;
        };
        let _nw104 = Array((new BigNumber(4)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw104.length); _i0_12++) {
          _nw104[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        (_742_v1)[_index139] = _nw104;
        let _745_v3;
        _745_v3 = _dafny.Seq.UnicodeFromString("rn");
        _745_v3 = _745_v3;
        if (!((_this).f19)) {
          let _746_v4;
          _746_v4 = new BigNumber(687);
          let _747_v5;
          _747_v5 = _dafny.Set.fromElements((_742_v1)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_742_v1).length))], _743_v2, (_742_v1)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_742_v1).length))], (_742_v1)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_742_v1).length))]);
          _746_v4 = new BigNumber(((_747_v5).Difference(_747_v5)).length);
          let _748_v6;
          let _init13 = ((_749_v3) => function (_750_i3) {
            return _749_v3;
          })(_745_v3);
          let _nw105 = Array((new BigNumber(21)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw105.length); _i0_13++) {
            _nw105[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _748_v6 = _nw105;
          let _751_v7;
          _751_v7 = new _dafny.CodePoint('n'.codePointAt(0));
          let _index140 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_748_v6).length));
          (_748_v6)[_index140] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), ((_752_v3, _753_p0) => function (_754_i4) {
            return (_752_v3)[_module.__default.safeIndex(_753_p0, new BigNumber((_752_v3).length))];
          })(_745_v3, p0)), _module.__default.safeIndex(_746_v4, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), ((_755_v3, _756_p0) => function (_757_i4) {
            return (_755_v3)[_module.__default.safeIndex(_756_p0, new BigNumber((_755_v3).length))];
          })(_745_v3, p0))).length)), _751_v7);
          let _758_v8;
          _758_v8 = _dafny.Seq.of(_745_v3);
          let _index141 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_748_v6).length));
          (_748_v6)[_index141] = (_758_v8)[_module.__default.safeIndex(p0, new BigNumber((_758_v8).length))];
          let _759_v9;
          _759_v9 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_739_i0),new BigNumber(-33));
          _746_v4 = (new BigNumber((_759_v9).length)).multipliedBy(_746_v4);
          let _760_v10;
          _760_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,p0);
          _760_v10 = (_760_v10).update(_this.f13, new BigNumber(-151));
          let _761_v11;
          let _nw106 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _761_v11 = _nw106;
          let _762_v12;
          let _nw107 = new _module.C1();
          _nw107.__ctor((_this).f19, _745_v3, _761_v11);
          _762_v12 = _nw107;
        } else {
          let _763_v13;
          _763_v13 = new BigNumber(-760);
          _763_v13 = (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f13,_763_v13)).length)).minus(new BigNumber(625)));
          let _764_v14;
          _764_v14 = new _dafny.CodePoint('j'.codePointAt(0));
          let _765_v15;
          _765_v15 = _dafny.Map.Empty.slice().updateUnsafe(_764_v14,(_this).f19);
          (_this).f13 = (((_765_v15).contains(_764_v14)) ? ((_765_v15).get(_764_v14)) : ((_this).f19));
          let _766_v16;
          let _init14 = ((_767_v13) => function (_768_i5) {
            return (_768_i5).plus((_dafny.ZERO).minus(_767_v13));
          })(_763_v13);
          let _nw108 = Array((new BigNumber(6)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw108.length); _i0_14++) {
            _nw108[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _766_v16 = _nw108;
          let _index142 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_766_v16).length));
          (_766_v16)[_index142] = new BigNumber(-945);
          let _769_v17;
          _769_v17 = _module.D0.create_DC1(_this.f13, p0);
          let _770_v18;
          let _nw109 = new _module.C2();
          _nw109.__ctor(_this.f13);
          _770_v18 = _nw109;
          let _771_v19;
          _771_v19 = _dafny.Seq.of(_770_v18);
          let _772_v20;
          _772_v20 = _dafny.Map.Empty.slice().updateUnsafe(_763_v13,_module.__default.fm28(new BigNumber((_771_v19).length), _763_v13, _this.f13, globalState));
          let _773_v21;
          let _nw110 = new _module.C0();
          _nw110.__ctor(new BigNumber(((((_772_v20).contains(_763_v13)) ? ((_772_v20).get(_763_v13)) : (_745_v3))).length));
          _773_v21 = _nw110;
          let _774_v22;
          _774_v22 = _dafny.Map.Empty.slice().updateUnsafe((_773_v21).f15,new BigNumber((_745_v3).length));
          let _775_v23;
          _775_v23 = _dafny.Set.fromElements((_this).fm4((_this).f19, _769_v17, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_773_v21,_770_v18.f13)).length), _774_v22, globalState));
          let _index143 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_766_v16).length));
          (_766_v16)[_index143] = new BigNumber((_775_v23).length);
          let _776_v24;
          _776_v24 = _module.D5.create_DC14(!(false));
          let _index144 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_766_v16).length));
          let _rhs107 = _776_v24;
          let _rhs108 = (_dafny.ZERO).minus(_739_i0);
          let _rhs109 = _this.f13;
          let _lhs66 = _766_v16;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_766_v16).length));
          let _lhs68 = _this;
          _776_v24 = _rhs107;
          _lhs66[_lhs67] = _rhs108;
          _lhs68.f13 = _rhs109;
          let _777_v25;
          _777_v25 = _dafny.Map.Empty.slice().updateUnsafe((_776_v24).dtor_cf23,(_766_v16)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_766_v16).length))]);
          let _778_v26;
          _778_v26 = _dafny.MultiSet.fromElements(_739_i0, _739_i0);
          _777_v25 = (_777_v25).update((_778_v26).IsSubsetOf(_778_v26), _763_v13);
        }
      }
      let _hi7 = p0;
      for (let _779_i6 = p0; _779_i6.isLessThan(_hi7); _779_i6 = _779_i6.plus(_dafny.ONE)) {
        (_this).f13 = _this.f13;
        let _780_v27;
        _780_v27 = _module.D2.create_DC8(_this.f13, _this.f13, _this.f13);
        let _781_v28;
        _781_v28 = _dafny.Seq.of(_780_v27, _780_v27);
        let _782_v29;
        _782_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,false);
        _781_v28 = _dafny.Seq.of(_780_v27, _module.D2.create_DC8((_this).f19, _this.f13, (((_782_v29).contains((_this).f19)) ? ((_782_v29).get((_this).f19)) : (_this.f13))), _780_v27);
        let _783_v30;
        _783_v30 = new BigNumber(328);
        _783_v30 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_783_v30, _783_v30));
        let _784_v31;
        _784_v31 = _dafny.Seq.UnicodeFromString("orhdrkah");
        let _785_v32;
        let _nw111 = Array((new BigNumber(2)).toNumber()).fill(false);
        _785_v32 = _nw111;
        let _786_v33;
        _786_v33 = _module.D6.create_DC17((_dafny.ZERO).minus(_779_i6), p0, _785_v32);
        let _787_v34;
        _787_v34 = _module.D6.create_DC18(_786_v33);
        let _788_v35;
        _788_v35 = _dafny.Map.Empty.slice().updateUnsafe(_787_v34,_this.f13);
        let _789_v36;
        _789_v36 = _module.D1.create_DC5(_784_v31, _785_v32, new BigNumber((_788_v35).length));
        let _pat_let_tv18 = _783_v30;
        _789_v36 = function (_pat_let21_0) {
          return function (_790_dt__update__tmp_h0) {
            return function (_pat_let22_0) {
              return function (_791_dt__update_hcf8_h0) {
                return function (_pat_let23_0) {
                  return function (_792_dt__update_hcf10_h0) {
                    return _module.D1.create_DC5(_791_dt__update_hcf8_h0, (_790_dt__update__tmp_h0).dtor_cf9, _792_dt__update_hcf10_h0);
                  }(_pat_let23_0);
                }(_pat_let_tv18);
              }(_pat_let22_0);
            }(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("f"), _module.__default.safeIndex(new BigNumber(694), new BigNumber((_dafny.Seq.UnicodeFromString("f")).length)), new _dafny.CodePoint('s'.codePointAt(0))));
          }(_pat_let21_0);
        }(_789_v36);
      }
      let _793_v37;
      _793_v37 = _module.D6.create_DC16(_module.__default.fm38(_this.f13, globalState));
      let _794_v38;
      _794_v38 = _module.D6.create_DC18(_793_v37);
      let _795_v39;
      _795_v39 = _module.D6.create_DC18(_793_v37);
      _795_v39 = _795_v39;
      let _796_v40;
      _796_v40 = new _dafny.CodePoint('q'.codePointAt(0));
      let _797_v41;
      _797_v41 = _module.D7.create_DC20(p0, _796_v40);
      let _798_v42;
      _798_v42 = _module.D0.create_DC1(!(_this.f13), p0);
      let _799_v43;
      _799_v43 = _dafny.Seq.of(p0);
      let _800_v44;
      _800_v44 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new _dafny.CodePoint('x'.codePointAt(0)));
      let _801_v45;
      let _nw112 = Array((new BigNumber(16)).toNumber());
      _nw112[(_dafny.ZERO).toNumber()] = _796_v40;
      _nw112[(_dafny.ONE).toNumber()] = _796_v40;
      _nw112[(new BigNumber(2)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(3)).toNumber()] = (_797_v41).dtor_cf32;
      _nw112[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
      _nw112[(new BigNumber(5)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(6)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(7)).toNumber()] = ((_this.f13) ? (_796_v40) : (_796_v40));
      _nw112[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
      _nw112[(new BigNumber(9)).toNumber()] = (((_this).fm4((_this).f19, _798_v42, p0, _module.__default.fm39(new BigNumber((_799_v43).length), globalState), globalState)) ? (_796_v40) : (_796_v40));
      _nw112[(new BigNumber(10)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(11)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(12)).toNumber()] = _796_v40;
      _nw112[(new BigNumber(13)).toNumber()] = ((false) ? (_796_v40) : (_796_v40));
      _nw112[(new BigNumber(14)).toNumber()] = ((_this.f13) ? (new _dafny.CodePoint('e'.codePointAt(0))) : (_796_v40));
      _nw112[(new BigNumber(15)).toNumber()] = (((_800_v44).contains(true)) ? ((_800_v44).get(true)) : (_796_v40));
      _801_v45 = _nw112;
      _801_v45 = _801_v45;
      let _802_v46;
      _802_v46 = _dafny.Seq.UnicodeFromString("hvovm");
      let _803_v47;
      _803_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("fls")).length),p0);
      let _804_v48;
      _804_v48 = _dafny.Seq.of(!((_this).f19));
      let _805_v49;
      _805_v49 = _dafny.Set.fromElements((_this).f19);
      let _806_v50;
      let _nw113 = Array((new BigNumber(6)).toNumber());
      _nw113[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_803_v47, p0, globalState);
      _nw113[(_dafny.ONE).toNumber()] = (_804_v48)[_module.__default.safeIndex(_module.__default.fm31(globalState), new BigNumber((_804_v48).length))];
      _nw113[(new BigNumber(2)).toNumber()] = (_804_v48)[_module.__default.safeIndex(new BigNumber((_805_v49).length), new BigNumber((_804_v48).length))];
      _nw113[(new BigNumber(3)).toNumber()] = _this.f13;
      _nw113[(new BigNumber(4)).toNumber()] = (_this).f19;
      _nw113[(new BigNumber(5)).toNumber()] = _this.f13;
      _806_v50 = _nw113;
      let _807_v51;
      _807_v51 = _module.D1.create_DC5(_dafny.Seq.Concat(_802_v46, _802_v46), _806_v50, p0);
      let _source15 = _807_v51;
      if (_source15.is_DC3) {
        let _808___mcc_h0 = (_source15).cf4;
        let _809_cf4 = _808___mcc_h0;
        let _810_v52;
        _810_v52 = _module.D1.create_DC4(_806_v50, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), ((_811_cf4) => function (_812_i7) {
  return _811_cf4;
})(_809_cf4))).length)), _796_v40);
        let _813_v53;
        _813_v53 = _dafny.Set.fromElements(_810_v52, _810_v52, _810_v52, _810_v52);
        let _814_v54;
        _814_v54 = _dafny.Seq.of(_813_v53);
        if (((_814_v54)[_module.__default.safeIndex(_809_cf4, new BigNumber((_814_v54).length))]).IsProperSubsetOf((_813_v53).Difference(_813_v53))) {
          let _815_v55;
          _815_v55 = _dafny.Map.Empty.slice().updateUnsafe(!(false),p0);
          _815_v55 = (_815_v55).update(!(!(_this.f13) || (_this.f13)), new BigNumber((_dafny.Seq.of(_805_v49, _805_v49)).length));
          let _816_v56;
          let _init15 = ((_817_cf4) => function (_818_i8) {
            return (_818_i8).minus(_817_cf4);
          })(_809_cf4);
          let _nw114 = Array((new BigNumber(7)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw114.length); _i0_15++) {
            _nw114[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _816_v56 = _nw114;
          let _819_v57;
          let _nw115 = new _module.C1();
          _nw115.__ctor(_this.f13, _802_v46, _816_v56);
          _819_v57 = _nw115;
          let _820_v58;
          _820_v58 = _dafny.MultiSet.fromElements((_797_v41).dtor_cf31);
          _809_cf4 = (((_820_v58).contains(_809_cf4)) ? ((_820_v58).get(_809_cf4)) : ((new BigNumber(-547)).multipliedBy(_809_cf4)));
          let _821_v59;
          let _nw116 = Array((new BigNumber(3)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), ((_822_v40) => function (_823_i9) {
            return _822_v40;
          })(_796_v40));
          _nw116[(_dafny.ONE).toNumber()] = _802_v46;
          _nw116[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("f");
          _821_v59 = _nw116;
          let _index145 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_821_v59).length));
          (_821_v59)[_index145] = _802_v46;
          let _index146 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_821_v59).length));
          (_821_v59)[_index146] = _dafny.Seq.update(_819_v57.f18, _module.__default.safeIndex(p0, new BigNumber((_819_v57.f18).length)), _796_v40);
          let _rhs110 = _dafny.Seq.UnicodeFromString("bplas");
          let _rhs111 = _819_v57;
          let _rhs112 = _799_v43;
          let _lhs69 = _819_v57;
          _lhs69.f18 = _rhs110;
          _819_v57 = _rhs111;
          _799_v43 = _rhs112;
        } else {
          let _824_v60;
          let _init16 = ((_825_cf4) => function (_826_i10) {
            return (_826_i10).multipliedBy(_825_cf4);
          })(_809_cf4);
          let _nw117 = Array((new BigNumber(2)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw117.length); _i0_16++) {
            _nw117[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _824_v60 = _nw117;
          let _index147 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length));
          (_824_v60)[_index147] = _module.__default.fm31(globalState);
          let _index148 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length));
          (_824_v60)[_index148] = p0;
          let _827_v61;
          _827_v61 = _dafny.Map.Empty.slice().updateUnsafe(_824_v60,_824_v60);
          _827_v61 = (_827_v61).update(_824_v60, ((_this.f13) ? (_824_v60) : (_824_v60)));
          let _828_v62;
          _828_v62 = _dafny.Set.fromElements((_824_v60)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length))]);
          let _829_v63;
          _829_v63 = _module.D2.create_DC7(_802_v46, _809_cf4, new BigNumber(-989), _this.f13);
          let _830_v64;
          _830_v64 = _dafny.MultiSet.fromElements((_828_v62).Difference(_dafny.Set.fromElements((_829_v63).dtor_cf13, (_824_v60)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length))], (_824_v60)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), ((_831_v40) => function (_832_i11) {
            return _831_v40;
          })(_796_v40))).length))), _828_v62);
          _830_v64 = _830_v64;
          let _index149 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_824_v60).length));
          (_824_v60)[_index149] = _809_cf4;
          _796_v40 = new _dafny.CodePoint('v'.codePointAt(0));
        }
        (_this).f13 = (p0).isLessThanOrEqualTo(_module.__default.fm31(globalState));
        let _833_v65;
        let _nw118 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _833_v65 = _nw118;
        let _index150 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_833_v65).length));
        (_833_v65)[_index150] = ((_this.f13) ? (_802_v46) : (_802_v46));
        let _index151 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_833_v65).length));
        (_833_v65)[_index151] = (_module.D2.create_DC7(_802_v46, new BigNumber(-656), new BigNumber((_799_v43).length), _this.f13)).dtor_cf12;
        let _834_v66;
        let _init17 = ((_835_p0, _836_cf4) => function (_837_i12) {
          return _dafny.Map.Empty.slice().updateUnsafe(_835_p0,_836_cf4);
        })(p0, _809_cf4);
        let _nw119 = Array((new BigNumber(24)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw119.length); _i0_17++) {
          _nw119[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _834_v66 = _nw119;
        _834_v66 = _834_v66;
      } else if (_source15.is_DC4) {
        let _838___mcc_h1 = (_source15).cf5;
        let _839___mcc_h2 = (_source15).cf6;
        let _840___mcc_h3 = (_source15).cf7;
        let _841_cf7 = _840___mcc_h3;
        let _842_cf6 = _839___mcc_h2;
        let _843_cf5 = _838___mcc_h1;
        _842_cf6 = _module.__default.safeModuloInt(p0, (((_this).f19) ? (_module.__default.fm31(globalState)) : (_842_cf6)));
        let _844_v67;
        let _init18 = function (_845_i13) {
          return (_845_i13).multipliedBy(new BigNumber(862));
        };
        let _nw120 = Array((new BigNumber(28)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw120.length); _i0_18++) {
          _nw120[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _844_v67 = _nw120;
        let _846_v68;
        _846_v68 = _dafny.Seq.of(_dafny.Seq.of(_module.__default.fm31(globalState), new BigNumber((_805_v49).length)), _dafny.Seq.of(_842_cf6, p0), _799_v43, _799_v43);
        let _847_v69;
        _847_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f13);
        let _848_v70;
        _848_v70 = _dafny.MultiSet.fromElements(_this.f13, (_this).f19);
        let _index152 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
        (_844_v67)[_index152] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((((_803_v47).contains(p0)) ? ((_803_v47).get(p0)) : (new BigNumber((_846_v68).length))))), new BigNumber(((_847_v69).update(new BigNumber((_848_v70).cardinality()), (_this).f19)).length));
        let _index153 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
        (_844_v67)[_index153] = _842_cf6;
        let _849_v71;
        _849_v71 = _dafny.Map.Empty.slice().updateUnsafe(_841_cf7,_806_v50);
        _849_v71 = (_849_v71).update(_841_cf7, _806_v50);
        let _850_v72;
        _850_v72 = _dafny.MultiSet.fromElements(_842_cf6);
        let _851_v73;
        _851_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
        if ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(((_851_v73).update(_this.f13, _this.f13)).length)), (_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))], _842_cf6)).IsSubsetOf(_850_v72)) {
          let _852_v74;
          _852_v74 = _module.D1.create_DC4(_806_v50, _842_cf6, _796_v40);
          let _853_v75;
          let _nw121 = new _module.C2();
          _nw121.__ctor(_this.f13);
          _853_v75 = _nw121;
          let _854_v76;
          _854_v76 = _dafny.Seq.of(_853_v75);
          let _index154 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
          let _rhs113 = _841_cf7;
          let _rhs114 = (_852_v74).dtor_cf6;
          let _rhs115 = ((new BigNumber(890)).multipliedBy((_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))])).isLessThan(new BigNumber(948));
          let _rhs116 = ((_850_v72).Difference(_module.__default.fm40(new BigNumber((_854_v76).length), _this.f13, _this.f13, (_this).f19, globalState))).Difference(_850_v72);
          let _rhs117 = (_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))];
          let _lhs70 = _this;
          let _lhs71 = _844_v67;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
          _841_cf7 = _rhs113;
          _842_cf6 = _rhs114;
          _lhs70.f13 = _rhs115;
          _850_v72 = _rhs116;
          _lhs71[_lhs72] = _rhs117;
          let _855_v77;
          let _nw122 = Array((new BigNumber(27)).toNumber()).fill([]);
          _855_v77 = _nw122;
          let _856_v78;
          _856_v78 = _dafny.Map.Empty.slice().updateUnsafe(_855_v77,_module.__default.fm41(p0, globalState));
          _856_v78 = (_856_v78).update(_855_v77, _module.D4.create_DC11(!((_this).f19)));
          let _index155 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
          (_844_v67)[_index155] = (_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))];
          let _857_v79;
          _857_v79 = _dafny.Map.Empty.slice().updateUnsafe(_850_v72,(new BigNumber(888)).minus(_842_cf6));
          let _858_v80;
          _858_v80 = _dafny.Set.fromElements((_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))]);
          let _index156 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
          (_844_v67)[_index156] = (((_857_v79).contains(_module.__default.fm40((_dafny.ZERO).minus((_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))]), false, _this.f13, _this.f13, globalState))) ? ((_857_v79).get(_module.__default.fm40((_dafny.ZERO).minus((_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))]), false, _this.f13, _this.f13, globalState))) : (new BigNumber((_858_v80).length)));
          let _859_v81;
          let _nw123 = new _module.C2();
          _nw123.__ctor(_this.f13);
          _859_v81 = _nw123;
        } else {
          let _index157 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length));
          (_844_v67)[_index157] = (_844_v67)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_844_v67).length))];
          (_this).f13 = (_804_v48)[_module.__default.safeIndex((_842_cf6).plus(new BigNumber((_802_v46).length)), new BigNumber((_804_v48).length))];
          (_this).f13 = (p0).isLessThan(_842_cf6);
          let _860_v83;
          _860_v83 = _dafny.Map.Empty.slice().updateUnsafe(_802_v46,_796_v40);
          let _861_v84;
          _861_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(150)), ((_862_v40) => function (_863_i14) {
            return _862_v40;
          })(_796_v40)),(_804_v48)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_804_v48).length))]);
          let _864_v85;
          let _out30;
          _out30 = _module.__default.m0((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_860_v83).Keys.Elements) {
              let _865_v82 = _compr_31;
              if ((_860_v83).contains(_865_v82)) {
                _coll31.push([_865_v82,true]);
              }
            }
            return _coll31;
          }()).Merge(_861_v84), globalState);
          _864_v85 = _out30;
          _842_cf6 = (_dafny.ZERO).minus(p0);
        }
      } else if (_source15.is_DC5) {
        let _866___mcc_h4 = (_source15).cf8;
        let _867___mcc_h5 = (_source15).cf9;
        let _868___mcc_h6 = (_source15).cf10;
        let _869_cf10 = _868___mcc_h6;
        let _870_cf9 = _867___mcc_h5;
        let _871_cf8 = _866___mcc_h4;
        (_this).f13 = !((_this).f19);
        let _872_v86;
        _872_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(808),_806_v50);
        let _873_v87;
        _873_v87 = _dafny.Set.fromElements(_870_cf9, _870_cf9, (((_872_v86).contains(_869_cf10)) ? ((_872_v86).get(_869_cf10)) : (_870_cf9)));
        _873_v87 = (_dafny.Set.fromElements(_806_v50, _806_v50)).Difference(_873_v87);
        _869_cf10 = _869_cf10;
        let _index158 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_806_v50).length));
        (_806_v50)[_index158] = ((_this).f19) || (_this.f13);
        let _index159 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_806_v50).length));
        let _rhs118 = ((!(_this.f13) || (false)) ? (_870_cf9) : (_806_v50));
        let _rhs119 = _module.__default.safeModuloInt(_869_cf10, _869_cf10);
        let _rhs120 = _this.f13;
        let _lhs73 = _806_v50;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_806_v50).length));
        _870_cf9 = _rhs118;
        _869_cf10 = _rhs119;
        _lhs73[_lhs74] = _rhs120;
      } else {
        let _874___mcc_h7 = (_source15).cf3;
        let _875_cf3 = _874___mcc_h7;
        let _876_v88;
        _876_v88 = new BigNumber(120);
        let _877_v89;
        _877_v89 = _dafny.Seq.of(_804_v48);
        let _878_v90;
        _878_v90 = (_877_v89)[_module.__default.safeIndex(p0, new BigNumber((_877_v89).length))];
        _876_v88 = new BigNumber(((_878_v90)).length);
        let _879_v91;
        let _nw124 = new _module.C0();
        _nw124.__ctor(_876_v88);
        _879_v91 = _nw124;
        let _880_v92;
        let _init19 = ((_881_v91) => function (_882_i15) {
          return (_882_i15).multipliedBy((_881_v91).f15);
        })(_879_v91);
        let _nw125 = Array((new BigNumber(6)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw125.length); _i0_19++) {
          _nw125[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _880_v92 = _nw125;
        let _883_v93;
        let _nw126 = new _module.C1();
        _nw126.__ctor(_this.f13, _dafny.Seq.Concat(_802_v46, _dafny.Seq.UnicodeFromString("uyflgki")), _880_v92);
        _883_v93 = _nw126;
        _876_v88 = (_879_v91).f15;
      }
      let _884_v94;
      _884_v94 = _module.D2.create_DC8(true, _this.f13, _this.f13);
      (_this).f13 = !(function (_source16) {
        if (_source16.is_DC7) {
          let _885___mcc_h8 = (_source16).cf12;
          let _886___mcc_h9 = (_source16).cf13;
          let _887___mcc_h10 = (_source16).cf14;
          let _888___mcc_h11 = (_source16).cf15;
          let _889_cf15 = _888___mcc_h11;
          let _890_cf14 = _887___mcc_h10;
          let _891_cf13 = _886___mcc_h9;
          let _892_cf12 = _885___mcc_h8;
          return _this.f13;
        } else if (_source16.is_DC8) {
          let _893___mcc_h12 = (_source16).cf16;
          let _894___mcc_h13 = (_source16).cf17;
          let _895___mcc_h14 = (_source16).cf18;
          let _896_cf18 = _895___mcc_h14;
          let _897_cf17 = _894___mcc_h13;
          let _898_cf16 = _893___mcc_h12;
          return _898_cf16;
        } else {
          let _899___mcc_h15 = (_source16).cf11;
          let _900_cf11 = _899___mcc_h15;
          return (_this).f19;
        }
      }(_884_v94));
      let _901_v95;
      let _nw127 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _901_v95 = _nw127;
      let _902_v96;
      _902_v96 = _module.D4.create_DC10(_901_v95);
      r0 = (_902_v96).dtor_cf20;
      return r0;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _903_i0;
      _903_i0 = _dafny.ZERO;
      L8: {
        while (false) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_903_i0)) {
              break L8;
            }
            _903_i0 = (_903_i0).plus(_dafny.ONE);
            let _904_v0;
            let _nw128 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _904_v0 = _nw128;
            let _index160 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_904_v0).length));
            (_904_v0)[_index160] = p2;
            let _index161 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_904_v0).length));
            (_904_v0)[_index161] = p2;
            let _index162 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_904_v0).length));
            (_904_v0)[_index162] = ((_this.f13) ? (p0) : (p2));
            let _905_v1;
            let _init20 = function (_906_i1) {
              return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f13, (_this).f19))).Union(_dafny.MultiSet.fromElements(_this.f13));
            };
            let _nw129 = Array((new BigNumber(26)).toNumber());
            for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw129.length); _i0_20++) {
              _nw129[_i0_20] = _init20(new BigNumber(_i0_20));
            }
            _905_v1 = _nw129;
            let _907_v2;
            _907_v2 = _dafny.MultiSet.fromElements((_this).f19, (_this).f19);
            let _index163 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_905_v1).length));
            (_905_v1)[_index163] = _907_v2;
            let _index164 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_905_v1).length));
            (_905_v1)[_index164] = (_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(_this.f13));
            let _908_v3;
            _908_v3 = _dafny.Seq.of(p2);
            let _909_v4;
            _909_v4 = _dafny.Map.Empty.slice().updateUnsafe(_908_v3,_908_v3);
            (_this).f13 = !(_909_v4).contains(_908_v3);
          }
        }
      }
      let _910_v6;
      _910_v6 = _dafny.Seq.of(new BigNumber(785));
      let _911_v7;
      _911_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new BigNumber((_910_v6).length));
      let _912_v8;
      _912_v8 = _dafny.MultiSet.fromElements(new BigNumber((_911_v7).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f19)).cardinality()));
      let _913_v9;
      _913_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      r2 = (_module.__default.fm0(function () {
        let _coll32 = new _dafny.Map();
        for (const _compr_32 of (_912_v8).Elements) {
          let _914_v5 = _compr_32;
          if ((_912_v8).contains(_914_v5)) {
            _coll32.push([_module.__default.safeDivisionInt(_914_v5, p2),p0]);
          }
        }
        return _coll32;
      }(), new BigNumber((_913_v9).length), globalState)) && ((_this).f19);
      let _915_i2;
      _915_i2 = _dafny.ZERO;
      L9: {
        while ((p0).isLessThan(p0)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_915_i2)) {
              break L9;
            }
            _915_i2 = (_915_i2).plus(_dafny.ONE);
            let _916_v10;
            _916_v10 = _module.D0.create_DC0(_this.f13);
            let _917_v11;
            _917_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_916_v10).dtor_cf0);
            let _918_v13;
            _918_v13 = _dafny.Seq.of(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-25)), function (_919_i3) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            }), p1);
            let _920_v15;
            _920_v15 = _dafny.Seq.of(_this.f13, _this.f13);
            let _921_v16;
            _921_v16 = _dafny.Seq.of(_920_v15);
            let _922_v17;
            let _nw130 = Array((new BigNumber(18)).toNumber());
            _nw130[(_dafny.ZERO).toNumber()] = p2;
            _nw130[(_dafny.ONE).toNumber()] = p0;
            _nw130[(new BigNumber(2)).toNumber()] = p2;
            _nw130[(new BigNumber(3)).toNumber()] = new BigNumber((function () {
              let _coll33 = new _dafny.Set();
              for (const _compr_33 of (function () {
                let _coll34 = new _dafny.Map();
                for (const _compr_34 of (_918_v13).Elements) {
                  let _923_v12 = _compr_34;
                  if (_dafny.Seq.contains(_918_v13, _923_v12)) {
                    _coll34.push([_923_v12,(_this).f19]);
                  }
                }
                return _coll34;
              }()).Keys.Elements) {
                let _924_v14 = _compr_33;
                if ((function () {
                  let _coll35 = new _dafny.Map();
                  for (const _compr_35 of (_918_v13).Elements) {
                    let _923_v12 = _compr_35;
                    if (_dafny.Seq.contains(_918_v13, _923_v12)) {
                      _coll35.push([_923_v12,(_this).f19]);
                    }
                  }
                  return _coll35;
                }()).contains(_924_v14)) {
                  _coll33.add(_924_v14);
                }
              }
              return _coll33;
            }()).length);
            _nw130[(new BigNumber(4)).toNumber()] = p0;
            _nw130[(new BigNumber(5)).toNumber()] = p0;
            _nw130[(new BigNumber(6)).toNumber()] = p0;
            _nw130[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p2);
            _nw130[(new BigNumber(8)).toNumber()] = p2;
            _nw130[(new BigNumber(9)).toNumber()] = p2;
            _nw130[(new BigNumber(10)).toNumber()] = p0;
            _nw130[(new BigNumber(11)).toNumber()] = new BigNumber((_921_v16).length);
            _nw130[(new BigNumber(12)).toNumber()] = p2;
            _nw130[(new BigNumber(13)).toNumber()] = p2;
            _nw130[(new BigNumber(14)).toNumber()] = p0;
            _nw130[(new BigNumber(15)).toNumber()] = p2;
            _nw130[(new BigNumber(16)).toNumber()] = p2;
            _nw130[(new BigNumber(17)).toNumber()] = p0;
            _922_v17 = _nw130;
            let _925_v18;
            let _nw131 = new _module.C1();
            _nw131.__ctor((((_917_v11).contains(p2)) ? ((_917_v11).get(p2)) : (_module.__default.fm0(_913_v9, p0, globalState))), _dafny.Seq.UnicodeFromString("wceigad"), _922_v17);
            _925_v18 = _nw131;
            let _926_v19;
            _926_v19 = new _dafny.CodePoint('p'.codePointAt(0));
            _926_v19 = _926_v19;
            (_this).f13 = (_925_v18).f17;
            let _927_v20;
            _927_v20 = _dafny.Map.Empty.slice().updateUnsafe(_920_v15,(_this).f19);
            _927_v20 = (_927_v20).update(_920_v15, (_this).f19);
          }
        }
      }
      let _928_v21;
      _928_v21 = _module.D2.create_DC6(_910_v6);
      let _pat_let_tv19 = _910_v6;
      let _pat_let_tv20 = p2;
      let _pat_let_tv21 = _910_v6;
      let _pat_let_tv22 = _912_v8;
      if (function (_source17) {
        if (_source17.is_DC7) {
          let _929___mcc_h0 = (_source17).cf12;
          let _930___mcc_h1 = (_source17).cf13;
          let _931___mcc_h2 = (_source17).cf14;
          let _932___mcc_h3 = (_source17).cf15;
          let _933_cf15 = _932___mcc_h3;
          let _934_cf14 = _931___mcc_h2;
          let _935_cf13 = _930___mcc_h1;
          let _936_cf12 = _929___mcc_h0;
          return _this.f13;
        } else if (_source17.is_DC8) {
          let _937___mcc_h4 = (_source17).cf16;
          let _938___mcc_h5 = (_source17).cf17;
          let _939___mcc_h6 = (_source17).cf18;
          let _940_cf18 = _939___mcc_h6;
          let _941_cf17 = _938___mcc_h5;
          let _942_cf16 = _937___mcc_h4;
          return (new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)))).length)).isLessThan((_pat_let_tv19)[_module.__default.safeIndex(_pat_let_tv20, new BigNumber((_pat_let_tv21).length))]);
        } else {
          let _943___mcc_h7 = (_source17).cf11;
          let _944_cf11 = _943___mcc_h7;
          return (new BigNumber((_pat_let_tv22).cardinality())).isLessThan(new BigNumber(75));
        }
      }(_928_v21)) {
        _913_v9 = (_913_v9).update(p0, new BigNumber(733));
        let _945_v22;
        _945_v22 = _dafny.Set.fromElements(p0, new BigNumber(150));
        let _946_v23;
        let _init21 = ((_947_p0) => function (_948_i4) {
          return (_948_i4).plus(_947_p0);
        })(p0);
        let _nw132 = Array((new BigNumber(12)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw132.length); _i0_21++) {
          _nw132[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _946_v23 = _nw132;
        let _949_v24;
        _949_v24 = _dafny.Map.Empty.slice().updateUnsafe(_946_v23,!((_this).f19));
        let _950_v25;
        _950_v25 = _dafny.Map.Empty.slice().updateUnsafe(_945_v22,(_949_v24).Merge(_949_v24));
        _950_v25 = (_950_v25).update((_dafny.Set.fromElements(p0)).Difference(_945_v22), _949_v24);
        let _951_v27;
        let _nw133 = new _module.C1();
        _nw133.__ctor(_module.__default.fm0(function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-765), new BigNumber(322))) {
            let _952_v26 = _compr_36;
            if (((new BigNumber(-765)).isLessThanOrEqualTo(_952_v26)) && ((_952_v26).isLessThan(new BigNumber(322)))) {
              _coll36.push([(_952_v26).plus(new BigNumber((_912_v8).cardinality())),p0]);
            }
          }
          return _coll36;
        }(), (_910_v6)[_module.__default.safeIndex(p2, new BigNumber((_910_v6).length))], globalState), p1, _946_v23);
        _951_v27 = _nw133;
        _951_v27 = _951_v27;
        let _953_v28;
        _953_v28 = _module.D4.create_DC11((_951_v27).f17);
        let _954_v29;
        _954_v29 = _dafny.MultiSet.fromElements(_953_v28);
        let _955_v30;
        _955_v30 = _dafny.Seq.of(_954_v29, _954_v29, _dafny.MultiSet.fromElements(_953_v28), _954_v29, _954_v29);
        let _956_v31;
        let _nw134 = Array((new BigNumber(12)).toNumber());
        _nw134[(_dafny.ZERO).toNumber()] = _this.f13;
        _nw134[(_dafny.ONE).toNumber()] = false;
        _nw134[(new BigNumber(2)).toNumber()] = (_951_v27).f17;
        _nw134[(new BigNumber(3)).toNumber()] = ((_951_v27).f17) || ((_this).f19);
        _nw134[(new BigNumber(4)).toNumber()] = (_this).f19;
        _nw134[(new BigNumber(5)).toNumber()] = _this.f13;
        _nw134[(new BigNumber(6)).toNumber()] = _this.f13;
        _nw134[(new BigNumber(7)).toNumber()] = _this.f13;
        _nw134[(new BigNumber(8)).toNumber()] = (_951_v27).f17;
        _nw134[(new BigNumber(9)).toNumber()] = _this.f13;
        _nw134[(new BigNumber(10)).toNumber()] = (_954_v29).IsDisjointFrom((_955_v30)[_module.__default.safeIndex(p2, new BigNumber((_955_v30).length))]);
        _nw134[(new BigNumber(11)).toNumber()] = (_951_v27).f17;
        _956_v31 = _nw134;
        let _957_v32;
        _957_v32 = _module.D5.create_DC14(!((_this).f19));
        let _index165 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_956_v31).length));
        (_956_v31)[_index165] = (function (_pat_let24_0) {
          return function (_958_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_959_dt__update_hcf23_h0) {
                return _module.D5.create_DC14(_959_dt__update_hcf23_h0);
              }(_pat_let25_0);
            }(_this.f13);
          }(_pat_let24_0);
        }(_957_v32)).dtor_cf23;
        let _index166 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_956_v31).length));
        (_956_v31)[_index166] = _this.f13;
        let _960_v33;
        _960_v33 = new BigNumber(470);
        _960_v33 = new BigNumber(203);
      } else {
        let _961_v34;
        _961_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
        let _rhs121 = !(!(!(!(true))) || (!((_this).f19)));
        let _rhs122 = !(!(((p2).isLessThanOrEqualTo(new BigNumber(((_961_v34).update(_this.f13, (_this).f19)).length))) || (_this.f13)));
        r2 = _rhs121;
        r1 = _rhs122;
        let _962_v35;
        _962_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.safeModuloInt(p0, p0));
        let _963_v36;
        _963_v36 = new _dafny.CodePoint('d'.codePointAt(0));
        _962_v35 = (_962_v35).update(_dafny.Seq.Concat(_dafny.Seq.update(p1, _module.__default.safeIndex(p0, new BigNumber((p1).length)), _963_v36), p1), (_dafny.ZERO).minus(p0));
        _963_v36 = _module.__default.fm34(globalState);
        let _964_v38;
        let _init22 = ((_965_v36) => function (_966_i5) {
          return (_966_i5).plus((_module.D7.create_DC20(new BigNumber((function () {
  let _coll37 = new _dafny.Map();
  for (const _compr_37 of _dafny.IntegerRange(new BigNumber(-622), new BigNumber(300))) {
    let _967_v37 = _compr_37;
    if (((new BigNumber(-622)).isLessThanOrEqualTo(_967_v37)) && ((_967_v37).isLessThan(new BigNumber(300)))) {
      _coll37.push([_module.__default.safeDivisionInt(_967_v37, new BigNumber((_dafny.Seq.of(true)).length)),new BigNumber(969)]);
    }
  }
  return _coll37;
}()).length), (_module.D7.create_DC20(new BigNumber((_dafny.Seq.of(true)).length), _965_v36)).dtor_cf32)).dtor_cf31);
        })(_963_v36);
        let _nw135 = Array((new BigNumber(3)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw135.length); _i0_22++) {
          _nw135[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _964_v38 = _nw135;
        let _968_v39;
        _968_v39 = _dafny.Set.fromElements(_964_v38, _964_v38);
        let _rhs123 = _this.f13;
        let _rhs124 = !((_968_v39).IsDisjointFrom((_968_v39).Difference(_968_v39)));
        r1 = _rhs123;
        r2 = _rhs124;
        let _969_v40;
        let _nw136 = Array((new BigNumber(18)).toNumber()).fill(_module.D7.Default());
        _969_v40 = _nw136;
        let _970_v41;
        let _nw137 = new _module.C1();
        _nw137.__ctor((_this).f19, p1, _964_v38);
        _970_v41 = _nw137;
        let _971_v42;
        _971_v42 = _module.D7.create_DC19(_970_v41);
        let _972_v43;
        _972_v43 = _module.D7.create_DC21(_971_v42);
        let _index167 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_969_v40).length));
        (_969_v40)[_index167] = _972_v43;
        let _index168 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_969_v40).length));
        (_969_v40)[_index168] = _972_v43;
      }
      let _973_v44;
      _973_v44 = _dafny.Seq.of(_this.f13);
      let _974_v45;
      _974_v45 = _dafny.Seq.of((_973_v44)[_module.__default.safeIndex(p0, new BigNumber((_973_v44).length))]);
      let _975_v46;
      let _nw138 = Array((new BigNumber(26)).toNumber());
      _nw138[(_dafny.ZERO).toNumber()] = (_this).f19;
      _nw138[(_dafny.ONE).toNumber()] = _this.f13;
      _nw138[(new BigNumber(2)).toNumber()] = !((_this).f19);
      _nw138[(new BigNumber(3)).toNumber()] = (_974_v45)[_module.__default.safeIndex(p0, new BigNumber((_974_v45).length))];
      _nw138[(new BigNumber(4)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(5)).toNumber()] = (_this).fm4((_this).f19, _module.__default.fm42(globalState), p0, _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm31(globalState)), globalState);
      _nw138[(new BigNumber(6)).toNumber()] = !((_this).f19);
      _nw138[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsPrefixOf(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), function (_976_i6) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
      _nw138[(new BigNumber(8)).toNumber()] = (_this).f19;
      _nw138[(new BigNumber(9)).toNumber()] = (_this).f19;
      _nw138[(new BigNumber(10)).toNumber()] = !(_this.f13);
      _nw138[(new BigNumber(11)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(12)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(13)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(14)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(15)).toNumber()] = (_this).f19;
      _nw138[(new BigNumber(16)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(17)).toNumber()] = _dafny.areEqual(_910_v6, _910_v6);
      _nw138[(new BigNumber(18)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(19)).toNumber()] = (_this).f19;
      _nw138[(new BigNumber(20)).toNumber()] = (_this).f19;
      _nw138[(new BigNumber(21)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(22)).toNumber()] = false;
      _nw138[(new BigNumber(23)).toNumber()] = _this.f13;
      _nw138[(new BigNumber(24)).toNumber()] = !((_this).f19);
      _nw138[(new BigNumber(25)).toNumber()] = !(_this.f13);
      _975_v46 = _nw138;
      _975_v46 = _975_v46;
      let _977_v47;
      let _nw139 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _977_v47 = _nw139;
      let _index169 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_977_v47).length));
      (_977_v47)[_index169] = p0;
      let _index170 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_977_v47).length));
      (_977_v47)[_index170] = _module.__default.fm31(globalState);
      let _978_v48;
      _978_v48 = _dafny.Map.Empty.slice().updateUnsafe((_977_v47)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_977_v47).length))],_975_v46);
      r0 = (_978_v48).Merge(_978_v48);
      r1 = (_dafny.MultiSet.fromElements(_this.f13, (_this).f19)).contains((true) || ((_this).f19));
      r2 = (_this).f19;
      r3 = _dafny.Seq.of(true, !((_977_v47)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_977_v47).length))]).isEqualTo(p2), (_this).f19, (((_this).f19) ? (_this.f13) : (true)));
      return [r0, r1, r2, r3];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f13 = false;
      this._f16 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    set f13(value) {
      let _this = this;
      _this._f13 = value;
    };
    __ctor(f16, f13) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f13 = f13;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f16;
    };
    fm5(globalState) {
      let _this = this;
      return _module.D1.create_DC3(new BigNumber(621));
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return true;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let _hi8 = p0;
      for (let _979_i0 = p0; _979_i0.isLessThan(_hi8); _979_i0 = _979_i0.plus(_dafny.ONE)) {
        let _980_v0;
        let _nw140 = new _module.C0();
        _nw140.__ctor(_979_i0);
        _980_v0 = _nw140;
        let _981_v1;
        let _init23 = function (_982_i1) {
          return ((_this.f13) ? (_dafny.Seq.UnicodeFromString("hwoxu")) : (_dafny.Seq.UnicodeFromString("pxu")));
        };
        let _nw141 = Array((new BigNumber(23)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw141.length); _i0_23++) {
          _nw141[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _981_v1 = _nw141;
        let _983_v2;
        _983_v2 = _dafny.Seq.UnicodeFromString("fylr");
        let _index171 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_981_v1).length));
        (_981_v1)[_index171] = _983_v2;
        let _index172 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_981_v1).length));
        (_981_v1)[_index172] = _983_v2;
        _981_v1 = _981_v1;
        let _984_v3;
        _984_v3 = _dafny.Seq.of((_980_v0).f15);
        let _985_v4;
        _985_v4 = _dafny.Seq.of(_this.f13);
        let _986_v5;
        _986_v5 = _module.D0.create_DC1(_this.f13, p0);
        let _987_v6;
        _987_v6 = new _dafny.CodePoint('k'.codePointAt(0));
        let _988_v7;
        _988_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f16);
        let _989_v8;
        let _nw142 = Array((new BigNumber(2)).toNumber());
        _nw142[(_dafny.ZERO).toNumber()] = _988_v7;
        _nw142[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13((_this).f16, (_this).f16, _987_v6, globalState),_this.f13);
        _989_v8 = _nw142;
        let _990_v9;
        _990_v9 = _dafny.Map.Empty.slice().updateUnsafe(_989_v8,(_980_v0).f15);
        let _991_v10;
        let _nw143 = Array((new BigNumber(9)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw143[(_dafny.ONE).toNumber()] = new BigNumber(((((_this).f16) ? (_985_v4) : (_985_v4))).length);
        _nw143[(new BigNumber(2)).toNumber()] = (_980_v0).f15;
        _nw143[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(p0, _module.__default.fm13((_986_v5).dtor_cf1, (_this).f16, _987_v6, globalState));
        _nw143[(new BigNumber(4)).toNumber()] = (((_990_v9).contains(_989_v8)) ? ((_990_v9).get(_989_v8)) : (p0));
        _nw143[(new BigNumber(5)).toNumber()] = new BigNumber(355);
        _nw143[(new BigNumber(6)).toNumber()] = (p0).plus((_980_v0).f15);
        _nw143[(new BigNumber(7)).toNumber()] = p0;
        _nw143[(new BigNumber(8)).toNumber()] = p0;
        _991_v10 = _nw143;
        let _index173 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_991_v10).length));
        (_991_v10)[_index173] = (_979_i0).minus(p0);
        let _992_v11;
        _992_v11 = new BigNumber(151);
        let _993_v12;
        let _nw144 = new _module.C1();
        _nw144.__ctor((_this).f16, _dafny.Seq.UnicodeFromString("luf"), _991_v10);
        _993_v12 = _nw144;
        let _994_v13;
        _994_v13 = _dafny.MultiSet.fromElements((_this).f16, (_this).f16);
        let _995_v14;
        _995_v14 = _dafny.Seq.of(_993_v12);
        let _996_v15;
        _996_v15 = _module.D2.create_DC7(_983_v2, _992_v11, new BigNumber(518), (_this).f16);
        let _index174 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_991_v10).length));
        let _rhs125 = _984_v3;
        let _rhs126 = (new BigNumber(947)).isLessThan(new BigNumber(560));
        let _rhs127 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_992_v11), _992_v11)).minus((_980_v0).f15);
        let _rhs128 = _module.__default.safeDivisionInt(p0, (_979_i0).plus((_993_v12).fm7((_980_v0).f15, new BigNumber((_994_v13).cardinality()), true, globalState)));
        let _rhs129 = (_995_v14)[_module.__default.safeIndex(((_980_v0).f15).minus((_996_v15).dtor_cf13), new BigNumber((_995_v14).length))];
        let _lhs75 = _this;
        let _lhs76 = _991_v10;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_991_v10).length));
        _984_v3 = _rhs125;
        _lhs75.f13 = _rhs126;
        _lhs76[_lhs77] = _rhs127;
        _992_v11 = _rhs128;
        _993_v12 = _rhs129;
      }
      let _997_v16;
      _997_v16 = new _dafny.CodePoint('n'.codePointAt(0));
      let _998_v17;
      _998_v17 = _dafny.Seq.of(_997_v16, _module.__default.fm22(_this.f13, _this.f13, globalState));
      let _999_v18;
      let _nw145 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _999_v18 = _nw145;
      let _1000_v19;
      let _nw146 = new _module.C1();
      _nw146.__ctor(true, _998_v17, _999_v18);
      _1000_v19 = _nw146;
      let _1001_v20;
      _1001_v20 = _dafny.Seq.of(_1000_v19);
      let _1002_v21;
      _1002_v21 = _module.D0.create_DC1(_this.f13, new BigNumber((_1001_v20).length));
      let _1003_v22;
      _1003_v22 = _module.D2.create_DC7(_998_v17, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_this.f13)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_1004_p0) => function (_1005_i2) {
  return _1004_p0;
})(p0))).length), !((_1002_v21).dtor_cf1));
      let _1006_v25;
      _1006_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(393),_1003_v22);
      let _1007_v27;
      _1007_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length),p0);
      let _1008_v28;
      let _nw147 = Array((new BigNumber(29)).toNumber());
      _nw147[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22);
      _nw147[(_dafny.ONE).toNumber()] = function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of (function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of _dafny.IntegerRange(new BigNumber(457), new BigNumber(35))) {
            let _1009_v24 = _compr_39;
            if (((new BigNumber(457)).isLessThanOrEqualTo(_1009_v24)) && ((_1009_v24).isLessThan(new BigNumber(35)))) {
              _coll39.push([(_1009_v24).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()),(_1000_v19).f17)).length)),p0]);
            }
          }
          return _coll39;
        }()).Keys.Elements) {
          let _1010_v23 = _compr_38;
          if ((function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(457), new BigNumber(35))) {
              let _1009_v24 = _compr_40;
              if (((new BigNumber(457)).isLessThanOrEqualTo(_1009_v24)) && ((_1009_v24).isLessThan(new BigNumber(35)))) {
                _coll40.push([(_1009_v24).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()),(_1000_v19).f17)).length)),p0]);
              }
            }
            return _coll40;
          }()).contains(_1010_v23)) {
            _coll38.push([_module.__default.safeModuloInt(_1010_v23, p0),_1003_v22]);
          }
        }
        return _coll38;
      }();
      _nw147[(new BigNumber(2)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22);
      _nw147[(new BigNumber(4)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(5)).toNumber()] = function () {
        let _coll41 = new _dafny.Map();
        for (const _compr_41 of _dafny.IntegerRange(new BigNumber(500), new BigNumber(-712))) {
          let _1011_v26 = _compr_41;
          if (((new BigNumber(500)).isLessThanOrEqualTo(_1011_v26)) && ((_1011_v26).isLessThan(new BigNumber(-712)))) {
            _coll41.push([_module.__default.safeModuloInt(_1011_v26, p0),_1003_v22]);
          }
        }
        return _coll41;
      }();
      _nw147[(new BigNumber(6)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(7)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(8)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(9)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22)).update(p0, _1003_v22);
      _nw147[(new BigNumber(11)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm23(new BigNumber((_1000_v19.f18).length), (_this).f16, globalState))).Merge(_1006_v25);
      _nw147[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22);
      _nw147[(new BigNumber(14)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(15)).toNumber()] = (_1006_v25).Merge(_1006_v25);
      _nw147[(new BigNumber(16)).toNumber()] = (_1006_v25).Merge(_1006_v25);
      _nw147[(new BigNumber(17)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(18)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(19)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(20)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(21)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(22)).toNumber()] = (_1006_v25).Merge(_1006_v25);
      _nw147[(new BigNumber(23)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22)).update(p0, _1003_v22);
      _nw147[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22);
      _nw147[(new BigNumber(25)).toNumber()] = (_1006_v25).update(p0, _module.D2.create_DC7(_1000_v19.f18, p0, new BigNumber((_dafny.Set.fromElements((_this).f16, (_this).fm4((_this).f16, _1002_v21, p0, _1007_v27, globalState))).length), (_this).f16));
      _nw147[(new BigNumber(26)).toNumber()] = _1006_v25;
      _nw147[(new BigNumber(27)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22)).Merge(_1006_v25);
      _nw147[(new BigNumber(28)).toNumber()] = _1006_v25;
      _1008_v28 = _nw147;
      let _index175 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1008_v28).length));
      (_1008_v28)[_index175] = _1006_v25;
      let _index176 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1008_v28).length));
      (_1008_v28)[_index176] = ((_1006_v25).Merge(_1006_v25)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1003_v22));
      let _1012_v29;
      _1012_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1003_v22).dtor_cf15,p0);
      (_this).f13 = (p0).isLessThanOrEqualTo(new BigNumber(((_1012_v29).Merge(_1012_v29)).length));
      let _1013_v30;
      _1013_v30 = new BigNumber(888);
      _1013_v30 = _1013_v30;
      _1013_v30 = (_module.__default.safeModuloInt(p0, p0)).plus(new BigNumber(-301));
      let _hi9 = p0;
      for (let _1014_i3 = _1013_v30; _1014_i3.isLessThan(_hi9); _1014_i3 = _1014_i3.plus(_dafny.ONE)) {
        let _1015_v31;
        _1015_v31 = _dafny.Seq.of(p0);
        (_this).f13 = ((_1015_v31)[_module.__default.safeIndex(_1013_v30, new BigNumber((_1015_v31).length))]).isLessThan(new BigNumber((_1012_v29).length));
        let _1016_v32;
        _1016_v32 = _dafny.Seq.of((_1000_v19).f17);
        let _rhs130 = _dafny.Seq.Concat(_1016_v32, _1016_v32);
        let _rhs131 = new BigNumber((_dafny.MultiSet.fromElements((((_1007_v27).contains(new BigNumber((_1007_v27).length))) ? ((_1007_v27).get(new BigNumber((_1007_v27).length))) : (new BigNumber((_1001_v20).length))), _1014_i3, p0)).cardinality());
        _1016_v32 = _rhs130;
        _1013_v30 = _rhs131;
        (_this).f13 = _this.f13;
        let _1017_v33;
        let _nw148 = new _module.C1();
        _nw148.__ctor((_1000_v19).fm14(p0, (_1000_v19).f17, _1014_i3, _1013_v30, globalState), _1000_v19.f18, _999_v18);
        _1017_v33 = _nw148;
      }
      r0 = _999_v18;
      return r0;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _1018_v0;
      _1018_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("mvnyliga"));
      let _hi10 = new BigNumber(((((_1018_v0).contains(p2)) ? ((_1018_v0).get(p2)) : (p1))).length);
      for (let _1019_i0 = p2; _1019_i0.isLessThan(_hi10); _1019_i0 = _1019_i0.plus(_dafny.ONE)) {
        let _1020_v1;
        _1020_v1 = new BigNumber(-188);
        let _1021_v2;
        _1021_v2 = _dafny.MultiSet.fromElements(_this.f13);
        let _rhs132 = new BigNumber((_1021_v2).cardinality());
        let _rhs133 = new BigNumber(((_module.__default.fm24(_1020_v1, _this.f13, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f13,_1019_i0))).length);
        let _rhs134 = (_this.f13) === ((_this).f16);
        let _rhs135 = _this.f13;
        let _lhs78 = _this;
        _1020_v1 = _rhs132;
        _1020_v1 = _rhs133;
        r1 = _rhs134;
        _lhs78.f13 = _rhs135;
        let _1022_v3;
        let _nw149 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1022_v3 = _nw149;
        _1022_v3 = _1022_v3;
        let _1023_v4;
        let _nw150 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1023_v4 = _nw150;
        let _1024_v5;
        let _nw151 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
        _1024_v5 = _nw151;
        let _1025_v6;
        _1025_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(_1019_i0));
        let _index177 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1024_v5).length));
        (_1024_v5)[_index177] = (_1025_v6).update(p0, p2);
        let _1026_v7;
        _1026_v7 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index178 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1024_v5).length));
        let _rhs136 = _1023_v4;
        let _rhs137 = (_1025_v6).update(_module.__default.fm13((_this).f16, true, _1026_v7, globalState), p2);
        let _rhs138 = !(_this.f13);
        let _rhs139 = new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.Concat(p1, p1))).length);
        let _rhs140 = _1020_v1;
        let _lhs79 = _1024_v5;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1024_v5).length));
        let _lhs81 = _this;
        _1023_v4 = _rhs136;
        _lhs79[_lhs80] = _rhs137;
        _lhs81.f13 = _rhs138;
        _1020_v1 = _rhs139;
        _1020_v1 = _rhs140;
        let _1027_v8;
        _1027_v8 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f16) === (!((_this).f16)),_dafny.Seq.update(p1, _module.__default.safeIndex(p2, new BigNumber((p1).length)), _1026_v7));
        let _1028_v9;
        _1028_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1019_i0)).cardinality()),_dafny.Set.fromElements(_1020_v1));
        let _1029_v10;
        let _init24 = ((_1030_p2) => function (_1031_i1) {
          return (_1031_i1).multipliedBy(_1030_p2);
        })(p2);
        let _nw152 = Array((new BigNumber(15)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw152.length); _i0_24++) {
          _nw152[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1029_v10 = _nw152;
        let _1032_v11;
        _1032_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1028_v9).length),_1029_v10);
        _1027_v8 = (_1027_v8).update((new BigNumber((_1032_v11).length)).isLessThan(new BigNumber(-440)), _dafny.Seq.UnicodeFromString("dct"));
      }
      let _1033_v12;
      let _nw153 = Array((new BigNumber(3)).toNumber()).fill(false);
      _1033_v12 = _nw153;
      let _index179 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length));
      (_1033_v12)[_index179] = false;
      let _1034_v13;
      _1034_v13 = new BigNumber(539);
      let _index180 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1033_v12).length));
      (_1033_v12)[_index180] = (_this).f16;
      let _1035_v14;
      _1035_v14 = _dafny.Set.fromElements(true, _this.f13);
      let _1036_v15;
      _1036_v15 = _dafny.Set.fromElements(_1034_v13);
      let _1037_v16;
      _1037_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_this.f13);
      let _index181 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length));
      let _index182 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1033_v12).length));
      let _rhs141 = _this.f13;
      let _rhs142 = _module.__default.safeDivisionInt(_1034_v13, new BigNumber(-663));
      let _rhs143 = !((_1035_v14).Union(_module.__default.fm25(_1036_v15, _1037_v16, p0, _this.f13, globalState))).equals((_1035_v14).Difference(_1035_v14));
      let _rhs144 = _this.f13;
      let _lhs82 = _1033_v12;
      let _lhs83 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length));
      let _lhs84 = _1033_v12;
      let _lhs85 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1033_v12).length));
      let _lhs86 = _this;
      _lhs82[_lhs83] = _rhs141;
      _1034_v13 = _rhs142;
      _lhs84[_lhs85] = _rhs143;
      _lhs86.f13 = _rhs144;
      _1034_v13 = new BigNumber(326);
      let _1038_v17;
      _1038_v17 = _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("arq"), _1033_v12, new BigNumber(764));
      let _1039_v18;
      _1039_v18 = _module.D4.create_DC11((_this).f16);
      let _1040_v19;
      _1040_v19 = new _dafny.CodePoint('m'.codePointAt(0));
      let _1041_v20;
      _1041_v20 = _dafny.MultiSet.fromElements(_1034_v13, new BigNumber((_1035_v14).length), p0);
      let _1042_v21;
      _1042_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(_1034_v13));
      let _1043_v22;
      let _init25 = function (_1044_i2) {
        return _module.__default.safeDivisionInt(_1044_i2, new BigNumber(836));
      };
      let _nw154 = Array((new BigNumber(25)).toNumber());
      for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw154.length); _i0_25++) {
        _nw154[_i0_25] = _init25(new BigNumber(_i0_25));
      }
      _1043_v22 = _nw154;
      let _1045_v23;
      _1045_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1034_v13);
      let _1046_v24;
      _1046_v24 = _module.D5.create_DC12(_1045_v23);
      let _1047_v25;
      _1047_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1043_v22,new BigNumber(((_1046_v24).dtor_cf22).length));
      let _1048_v26;
      _1048_v26 = _dafny.Seq.of((_this).f16);
      let _1049_v27;
      _1049_v27 = _dafny.Seq.of(_1034_v13, p0);
      let _1050_v28;
      let _nw155 = Array((new BigNumber(26)).toNumber());
      _nw155[(_dafny.ZERO).toNumber()] = (_1038_v17).dtor_cf10;
      _nw155[(_dafny.ONE).toNumber()] = new BigNumber(168);
      _nw155[(new BigNumber(2)).toNumber()] = p0;
      _nw155[(new BigNumber(3)).toNumber()] = (new BigNumber((p1).length)).plus(new BigNumber(590));
      _nw155[(new BigNumber(4)).toNumber()] = p0;
      _nw155[(new BigNumber(5)).toNumber()] = _module.__default.fm13(!((_1039_v18).dtor_cf21), _this.f13, _1040_v19, globalState);
      _nw155[(new BigNumber(6)).toNumber()] = p0;
      _nw155[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_1041_v20).cardinality())), new BigNumber((_1042_v21).length)));
      _nw155[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(_module.__default.fm13((_this).f16, _this.f13, new _dafny.CodePoint('d'.codePointAt(0)), globalState), (_dafny.ZERO).minus(p0));
      _nw155[(new BigNumber(9)).toNumber()] = p0;
      _nw155[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p2, new BigNumber((_module.__default.fm26(_1040_v19, p0, _1034_v13, globalState)).length)));
      _nw155[(new BigNumber(11)).toNumber()] = p0;
      _nw155[(new BigNumber(12)).toNumber()] = (((_1047_v25).contains(_1043_v22)) ? ((_1047_v25).get(_1043_v22)) : (new BigNumber((p1).length)));
      _nw155[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(p2, new BigNumber((_1048_v26).length));
      _nw155[(new BigNumber(14)).toNumber()] = p2;
      _nw155[(new BigNumber(15)).toNumber()] = p0;
      _nw155[(new BigNumber(16)).toNumber()] = new BigNumber(532);
      _nw155[(new BigNumber(17)).toNumber()] = p2;
      _nw155[(new BigNumber(18)).toNumber()] = p0;
      _nw155[(new BigNumber(19)).toNumber()] = (_1034_v13).plus(_1034_v13);
      _nw155[(new BigNumber(20)).toNumber()] = (new BigNumber(707)).multipliedBy(p2);
      _nw155[(new BigNumber(21)).toNumber()] = (_1034_v13).plus(_1034_v13);
      _nw155[(new BigNumber(22)).toNumber()] = p2;
      _nw155[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(((_1049_v27)[_module.__default.safeIndex(p2, new BigNumber((_1049_v27).length))]).plus(p2));
      _nw155[(new BigNumber(24)).toNumber()] = new BigNumber(333);
      _nw155[(new BigNumber(25)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(_1040_v19, _1040_v19)).update(_1040_v19, _module.__default.abs(_1034_v13))).cardinality());
      _1050_v28 = _nw155;
      let _index183 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length));
      (_1043_v22)[_index183] = (p2).multipliedBy(new BigNumber((_1049_v27).length));
      let _1051_v29;
      _1051_v29 = _1037_v16;
      let _index184 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length));
      let _rhs145 = p0;
      let _rhs146 = _1051_v29;
      let _lhs87 = _1043_v22;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length));
      _lhs87[_lhs88] = _rhs145;
      _1051_v29 = _rhs146;
      let _1052_v30;
      let _nw156 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
      _1052_v30 = _nw156;
      let _1053_v31;
      _1053_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v26,(_1045_v23).update((_1033_v12)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length))], _1034_v13));
      let _index185 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1052_v30).length));
      (_1052_v30)[_index185] = _1053_v31;
      let _1054_v32;
      _1054_v32 = _module.D6.create_DC16(_1053_v31);
      let _1055_v33;
      _1055_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1033_v12)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length))],p1);
      let _index186 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length));
      let _index187 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1052_v30).length));
      let _index188 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length));
      let _rhs147 = !(_this.f13) || ((_1048_v26)[_module.__default.safeIndex(p2, new BigNumber((_1048_v26).length))]);
      let _rhs148 = (_1054_v32).dtor_cf25;
      let _rhs149 = (_dafny.ZERO).minus((p0).minus(((_1043_v22)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length))]).plus(new BigNumber((_1055_v33).length))));
      let _rhs150 = (p2).plus(new BigNumber(534));
      let _rhs151 = _1050_v28;
      let _lhs89 = _1033_v12;
      let _lhs90 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length));
      let _lhs91 = _1052_v30;
      let _lhs92 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1052_v30).length));
      let _lhs93 = _1043_v22;
      let _lhs94 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1043_v22).length));
      _lhs89[_lhs90] = _rhs147;
      _lhs91[_lhs92] = _rhs148;
      _lhs93[_lhs94] = _rhs149;
      _1034_v13 = _rhs150;
      _1043_v22 = _rhs151;
      _1035_v14 = _1035_v14;
      let _1056_v34;
      _1056_v34 = _dafny.Seq.of(_1033_v12);
      let _1057_v35;
      _1057_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1037_v16).length),(_1056_v34)[_module.__default.safeIndex(_1034_v13, new BigNumber((_1056_v34).length))]);
      r0 = (((_1033_v12)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length))]) ? (_1057_v35) : (_1057_v35));
      r1 = ((_this.f13) ? ((_1048_v26)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1048_v26)).cardinality()), new BigNumber((_1048_v26).length))]) : ((_1033_v12)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1033_v12).length))]));
      let _1058_v36;
      _1058_v36 = _dafny.MultiSet.fromElements((_this).f16, _this.f13, !((_this).f16));
      r2 = ((_1058_v36).Difference(_1058_v36)).IsSubsetOf(_1058_v36);
      r3 = _1048_v26;
      return [r0, r1, r2, r3];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1059_v0;
      _1059_v0 = _module.D4.create_DC11(p1);
      let _1060_v1;
      _1060_v1 = new _dafny.CodePoint('o'.codePointAt(0));
      let _1061_v2;
      _1061_v2 = _dafny.Seq.of(_1060_v1);
      let _1062_v3;
      _1062_v3 = new BigNumber(956);
      let _1063_v4;
      _1063_v4 = _module.D2.create_DC7(_1061_v2, _1062_v3, new BigNumber(593), p0);
      let _pat_let_tv23 = _1063_v4;
      let _source18 = function (_pat_let26_0) {
        return function (_1064_dt__update__tmp_h0) {
          return function (_pat_let27_0) {
            return function (_1065_dt__update_hcf21_h0) {
              return _module.D4.create_DC11(_1065_dt__update_hcf21_h0);
            }(_pat_let27_0);
          }((_pat_let_tv23).dtor_cf15);
        }(_pat_let26_0);
      }(_1059_v0);
      if (_source18.is_DC11) {
        let _1066___mcc_h0 = (_source18).cf21;
        let _1067_cf21 = _1066___mcc_h0;
        let _1068_v5;
        let _nw157 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _1068_v5 = _nw157;
        let _index189 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1068_v5).length));
        (_1068_v5)[_index189] = _1062_v3;
        let _index190 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1068_v5).length));
        (_1068_v5)[_index190] = _module.__default.safeModuloInt(_module.__default.fm13(_1067_cf21, _1067_cf21, _1060_v1, globalState), (_dafny.ZERO).minus((_dafny.ZERO).minus((_1062_v3).multipliedBy(_1062_v3))));
        let _1069_v6;
        _1069_v6 = _dafny.Set.fromElements(p1, _this.f13, _1067_cf21);
        _1069_v6 = _1069_v6;
        let _1070_v7;
        _1070_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,!(p1));
        _1070_v7 = (_1070_v7).update(p1, _1067_cf21);
        let _1071_v8;
        let _nw158 = new _module.C1();
        _nw158.__ctor(_1067_cf21, _dafny.Seq.update(_1061_v2, _module.__default.safeIndex(new BigNumber((_1070_v7).length), new BigNumber((_1061_v2).length)), new _dafny.CodePoint('c'.codePointAt(0))), _1068_v5);
        _1071_v8 = _nw158;
      } else {
        let _1072___mcc_h1 = (_source18).cf20;
        let _1073_cf20 = _1072___mcc_h1;
        let _1074_v9;
        let _init26 = ((_1075_p1) => function (_1076_i0) {
          return _1075_p1;
        })(p1);
        let _nw159 = Array((new BigNumber(10)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw159.length); _i0_26++) {
          _nw159[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1074_v9 = _nw159;
        let _1077_v10;
        _1077_v10 = _module.D1.create_DC4(_1074_v9, _1062_v3, _module.__default.fm22((_this).f16, p0, globalState));
        let _1078_v11;
        let _nw160 = Array((new BigNumber(28)).toNumber());
        _nw160[(_dafny.ZERO).toNumber()] = _1060_v1;
        _nw160[(_dafny.ONE).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(2)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(3)).toNumber()] = (_1077_v10).dtor_cf7;
        _nw160[(new BigNumber(4)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(5)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(6)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw160[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw160[(new BigNumber(9)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(10)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw160[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw160[(new BigNumber(13)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(14)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw160[(new BigNumber(16)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(17)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(18)).toNumber()] = (_1061_v2)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1061_v2).length))];
        _nw160[(new BigNumber(19)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(20)).toNumber()] = (_1061_v2)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1061_v2).length))];
        _nw160[(new BigNumber(21)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(22)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(23)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(24)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(25)).toNumber()] = _1060_v1;
        _nw160[(new BigNumber(26)).toNumber()] = (_1077_v10).dtor_cf7;
        _nw160[(new BigNumber(27)).toNumber()] = _1060_v1;
        _1078_v11 = _nw160;
        let _index191 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1078_v11).length));
        (_1078_v11)[_index191] = _module.__default.fm22(_this.f13, p0, globalState);
        let _index192 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1078_v11).length));
        let _rhs152 = (!(((_this.f13) ? ((_this).f16) : (p1)))) || ((p1) && ((_this).f16));
        let _rhs153 = _1060_v1;
        let _lhs95 = _this;
        let _lhs96 = _1078_v11;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1078_v11).length));
        _lhs95.f13 = _rhs152;
        _lhs96[_lhs97] = _rhs153;
        let _1079_v12;
        _1079_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v3,_1062_v3);
        let _1080_v13;
        let _nw161 = new _module.C2();
        _nw161.__ctor(true);
        _1080_v13 = _nw161;
        r1 = (new BigNumber(459)).isLessThan((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1079_v12,_1080_v13)).length)).plus(new BigNumber((_1061_v2).length)));
        let _1081_v14;
        _1081_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v3,(_this).f16);
        let _1082_v15;
        _1082_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1081_v14).contains((_dafny.ZERO).minus(_1062_v3))) ? ((_1081_v14).get((_dafny.ZERO).minus(_1062_v3))) : (p0)),_module.__default.safeDivisionInt(_1062_v3, _1062_v3));
        _1082_v15 = (_1082_v15).update(_1080_v13.f13, _module.__default.safeDivisionInt(_1062_v3, _1062_v3));
        _1081_v14 = _1081_v14;
      }
      (_this).f13 = _this.f13;
      let _1083_v16;
      _1083_v16 = _dafny.Seq.of(!(_this.f13), p1, _this.f13, (_this).f16);
      let _1084_v17;
      let _init27 = function (_1085_i1) {
        return _this.f13;
      };
      let _nw162 = Array((new BigNumber(11)).toNumber());
      for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw162.length); _i0_27++) {
        _nw162[_i0_27] = _init27(new BigNumber(_i0_27));
      }
      _1084_v17 = _nw162;
      let _1086_v18;
      _1086_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v16,_1084_v17);
      _1086_v18 = (_1086_v18).update(_dafny.Seq.of((_1083_v16)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1083_v16).length))], (_this).f16, !((_this).f16)), _1084_v17);
      let _1087_i2;
      _1087_i2 = _dafny.ZERO;
      L10: {
        while (p1) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1087_i2)) {
              break L10;
            }
            _1087_i2 = (_1087_i2).plus(_dafny.ONE);
            if ((_this).f16) {
              let _1088_v19;
              _1088_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_1062_v3);
              let _1089_v20;
              _1089_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_1088_v19).length));
              let _1090_v21;
              _1090_v21 = _dafny.Set.fromElements(true, !(p0));
              let _1091_v22;
              _1091_v22 = _dafny.Set.fromElements(_1062_v3, new BigNumber((_1090_v21).length));
              let _1092_v23;
              _1092_v23 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,p0);
              let _rhs154 = (_1083_v16)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1062_v3, _1062_v3), new BigNumber((_1083_v16).length))];
              let _rhs155 = (((p0) ? (new BigNumber((_1088_v19).length)) : (_1062_v3))).multipliedBy(_module.__default.safeDivisionInt(_1062_v3, _1062_v3));
              let _rhs156 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1062_v3, new BigNumber(-641), new BigNumber((_1061_v2).length), _1062_v3),_module.__default.fm25(_1091_v22, _1092_v23, new BigNumber((_dafny.Seq.of(_this.f13, false, (_this).f16)).length), p0, globalState))).length);
              r0 = _rhs154;
              _1062_v3 = _rhs155;
              r2 = _rhs156;
              let _1093_v24;
              _1093_v24 = _module.D0.create_DC1((_this).f16, new BigNumber(-634));
              let _1094_v25;
              _1094_v25 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1062_v3),_1062_v3);
              let _1095_v26;
              _1095_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1061_v2).length),new BigNumber((_1094_v25).length));
              let _1096_v27;
              _1096_v27 = _dafny.Seq.of(_1062_v3);
              _1060_v1 = _module.__default.fm22((_this).fm4((_1083_v16)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1083_v16).length))], _1093_v24, _1062_v3, (_1094_v25).update(_1062_v3, _1062_v3), globalState), ((_1096_v27)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1096_v27).length))]).isLessThanOrEqualTo(_1062_v3), globalState);
              let _1097_v28;
              let _nw163 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
              _1097_v28 = _nw163;
              let _index193 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1097_v28).length));
              (_1097_v28)[_index193] = _1062_v3;
              let _index194 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1097_v28).length));
              let _rhs157 = ((new BigNumber((_1061_v2).length)).minus(_1062_v3)).isLessThan(new BigNumber((_1089_v20).length));
              let _rhs158 = _this.f13;
              let _rhs159 = _1062_v3;
              let _lhs98 = _this;
              let _lhs99 = _1097_v28;
              let _lhs100 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1097_v28).length));
              _lhs98.f13 = _rhs157;
              r0 = _rhs158;
              _lhs99[_lhs100] = _rhs159;
              _1096_v27 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), ((_1098_v28) => function (_1099_i3) {
                return (_1098_v28)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_1098_v28).length))];
              })(_1097_v28));
              let _1100_v29;
              _1100_v29 = _dafny.MultiSet.fromElements(((_1097_v28)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_1097_v28).length))]).minus(_1062_v3), new BigNumber(741), (_1097_v28)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_1097_v28).length))], _1062_v3);
              let _1101_v30;
              let _nw164 = new _module.C1();
              _nw164.__ctor(p0, _dafny.Seq.UnicodeFromString("nldyh"), _1097_v28);
              _1101_v30 = _nw164;
              let _1102_v31;
              _1102_v31 = _dafny.Seq.of(_1101_v30, _1101_v30);
              let _1103_v32;
              _1103_v32 = _module.D7.create_DC19(_1101_v30);
              let _1104_v33;
              let _nw165 = Array((new BigNumber(14)).toNumber());
              _nw165[(_dafny.ZERO).toNumber()] = _1101_v30;
              _nw165[(_dafny.ONE).toNumber()] = (_1102_v31)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1102_v31).length))];
              _nw165[(new BigNumber(2)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(3)).toNumber()] = (_1103_v32).dtor_cf30;
              _nw165[(new BigNumber(4)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(5)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(6)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(7)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(8)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(9)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(10)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(11)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(12)).toNumber()] = _1101_v30;
              _nw165[(new BigNumber(13)).toNumber()] = _1101_v30;
              _1104_v33 = _nw165;
              let _1105_v34;
              _1105_v34 = _dafny.Seq.of(_1104_v33);
              let _rhs160 = _1060_v1;
              let _rhs161 = ((p1) ? (_1083_v16) : (_1083_v16));
              let _rhs162 = _dafny.Seq.IsPrefixOf(_1061_v2, _1061_v2);
              let _rhs163 = (_1100_v29).update(_1062_v3, _module.__default.abs(new BigNumber((_1090_v21).length)));
              let _rhs164 = _1105_v34;
              _1060_v1 = _rhs160;
              _1083_v16 = _rhs161;
              r1 = _rhs162;
              _1100_v29 = _rhs163;
              _1105_v34 = _rhs164;
            } else {
              let _1106_v35;
              let _nw166 = Array((new BigNumber(4)).toNumber());
              _nw166[(_dafny.ZERO).toNumber()] = _1062_v3;
              _nw166[(_dafny.ONE).toNumber()] = new BigNumber(330);
              _nw166[(new BigNumber(2)).toNumber()] = _1062_v3;
              _nw166[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_1062_v3, _1062_v3);
              _1106_v35 = _nw166;
              let _1107_v36;
              _1107_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1106_v35,p0);
              let _index195 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length));
              (_1106_v35)[_index195] = new BigNumber((_1107_v36).length);
              let _1108_v37;
              _1108_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,(_1083_v16)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1083_v16).length))]);
              let _1109_v38;
              _1109_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1108_v37,_1062_v3);
              let _index196 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length));
              (_1106_v35)[_index196] = new BigNumber(((_1109_v38).update(_1108_v37, _1062_v3)).length);
              let _1110_v39;
              let _nw167 = new _module.C1();
              _nw167.__ctor(p0, _1061_v2, _1106_v35);
              _1110_v39 = _nw167;
              let _1111_v40;
              _1111_v40 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f13),new BigNumber((_1061_v2).length));
              let _1112_v41;
              _1112_v41 = _dafny.Seq.of(_1062_v3, new BigNumber(431));
              let _1113_v42;
              _1113_v42 = _dafny.MultiSet.fromElements(new BigNumber((_1110_v39.f18).length));
              let _1114_v43;
              _1114_v43 = _dafny.Seq.of((_1112_v41)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1112_v41).length))], (((_1113_v42).contains(new BigNumber((_1061_v2).length))) ? ((_1113_v42).get(new BigNumber((_1061_v2).length))) : ((_1106_v35)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length))])), (_1106_v35)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length))], _1062_v3);
              let _rhs165 = _dafny.Seq.Concat(((!((_this).f16)) ? (_1061_v2) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("aiwrcowth"), _module.__default.safeIndex(new BigNumber((_1111_v40).length), new BigNumber((_dafny.Seq.UnicodeFromString("aiwrcowth")).length)), _1060_v1))), _1110_v39.f18);
              let _rhs166 = !(!(p1));
              let _rhs167 = (_1114_v43)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1114_v43).length))];
              _1061_v2 = _rhs165;
              r0 = _rhs166;
              _1062_v3 = _rhs167;
              r2 = (_1106_v35)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length))];
              _1113_v42 = (_dafny.MultiSet.FromArray(_1112_v41)).update(_1062_v3, _module.__default.abs(((_1106_v35)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1106_v35).length))]).minus(new BigNumber(31))));
            }
            let _1115_v44;
            _1115_v44 = _module.D1.create_DC4(_1084_v17, _1062_v3, _1060_v1);
            let _1116_v45;
            _1116_v45 = _dafny.Seq.of(_1061_v2, _dafny.Seq.update(_1061_v2, _module.__default.safeIndex(_1062_v3, new BigNumber((_1061_v2).length)), new _dafny.CodePoint('l'.codePointAt(0))), _1061_v2);
            let _1117_v46;
            let _nw168 = Array((new BigNumber(11)).toNumber());
            _nw168[(_dafny.ZERO).toNumber()] = (_1062_v3).minus((_1115_v44).dtor_cf6);
            _nw168[(_dafny.ONE).toNumber()] = new BigNumber(609);
            _nw168[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_1062_v3, new BigNumber((_dafny.Seq.of((_1063_v4).dtor_cf13)).length));
            _nw168[(new BigNumber(3)).toNumber()] = _1062_v3;
            _nw168[(new BigNumber(4)).toNumber()] = _1062_v3;
            _nw168[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_1062_v3);
            _nw168[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_1062_v3, new BigNumber((_1083_v16).length));
            _nw168[(new BigNumber(7)).toNumber()] = (_1062_v3).multipliedBy(_1062_v3);
            _nw168[(new BigNumber(8)).toNumber()] = (_1062_v3).multipliedBy(_1062_v3);
            _nw168[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_1062_v3).multipliedBy(new BigNumber((_1116_v45).length)));
            _nw168[(new BigNumber(10)).toNumber()] = _1062_v3;
            _1117_v46 = _nw168;
            let _index197 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
            (_1117_v46)[_index197] = new BigNumber((_1061_v2).length);
            let _1118_v47;
            _1118_v47 = _module.D6.create_DC17(_1062_v3, _1062_v3, _1084_v17);
            let _1119_v48;
            _1119_v48 = _dafny.Set.fromElements(_1118_v47);
            let _index198 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
            (_1117_v46)[_index198] = new BigNumber(((_1119_v48).Union(_1119_v48)).length);
            let _1120_v49;
            _1120_v49 = _dafny.MultiSet.fromElements(_1060_v1);
            let _1121_v50;
            _1121_v50 = _dafny.Seq.of(_module.__default.fm31(globalState));
            let _rhs168 = new BigNumber(120);
            let _rhs169 = new BigNumber((_dafny.Seq.UnicodeFromString("dcdkyipa")).length);
            let _rhs170 = _dafny.MultiSet.FromArray(_module.__default.fm28(((_this.f13) ? (new BigNumber((_1061_v2).length)) : ((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])), _module.__default.fm31(globalState), !_dafny.areEqual(_1121_v50, _1121_v50), globalState));
            let _rhs171 = _1117_v46;
            _1062_v3 = _rhs168;
            _1062_v3 = _rhs169;
            _1120_v49 = _rhs170;
            _1117_v46 = _rhs171;
            let _1122_v51;
            _1122_v51 = _dafny.Seq.of(_1121_v50);
            let _1123_v52;
            _1123_v52 = _module.D2.create_DC6((_1122_v51)[_module.__default.safeIndex(_module.__default.fm31(globalState), new BigNumber((_1122_v51).length))]);
            let _source19 = _1123_v52;
            if (_source19.is_DC7) {
              let _1124___mcc_h2 = (_source19).cf12;
              let _1125___mcc_h3 = (_source19).cf13;
              let _1126___mcc_h4 = (_source19).cf14;
              let _1127___mcc_h5 = (_source19).cf15;
              let _1128_cf15 = _1127___mcc_h5;
              let _1129_cf14 = _1126___mcc_h4;
              let _1130_cf13 = _1125___mcc_h3;
              let _1131_cf12 = _1124___mcc_h2;
              let _1132_v53;
              _1132_v53 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ajufne"), _1061_v2),(_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))]);
              _1132_v53 = _1132_v53;
              let _index199 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1084_v17).length));
              (_1084_v17)[_index199] = _module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))],(_dafny.ZERO).minus(_1062_v3)), (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))], globalState);
              let _1133_v54;
              _1133_v54 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1120_v49).cardinality())),_this.f13);
              let _1134_v55;
              _1134_v55 = _dafny.Set.fromElements(_1083_v16);
              let _1135_v56;
              let _nw169 = new _module.C2();
              _nw169.__ctor(true);
              _1135_v56 = _nw169;
              let _1136_v57;
              _1136_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1135_v56,(_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))]);
              let _index200 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              let _index201 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1084_v17).length));
              let _rhs172 = (_1128_cf15) && (p1);
              let _rhs173 = _module.__default.safeModuloInt(_1130_cf13, (_1130_cf13).minus((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))]));
              let _rhs174 = !_dafny.Seq.contains(_1121_v50, (_1062_v3).minus(new BigNumber(635)));
              let _rhs175 = ((_1134_v55).Intersect(_1134_v55)).IsProperSubsetOf(_dafny.Set.fromElements(_1083_v16, _1083_v16));
              let _rhs176 = (_1133_v54).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1122_v51)[_module.__default.safeIndex((((_1136_v57).contains(_1135_v56)) ? ((_1136_v57).get(_1135_v56)) : ((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])), new BigNumber((_1122_v51).length))]).length),p0)).Merge(_1133_v54));
              let _lhs101 = _1117_v46;
              let _lhs102 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              let _lhs103 = _1084_v17;
              let _lhs104 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1084_v17).length));
              r1 = _rhs172;
              _lhs101[_lhs102] = _rhs173;
              _lhs103[_lhs104] = _rhs174;
              r0 = _rhs175;
              _1133_v54 = _rhs176;
              let _index202 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              (_1117_v46)[_index202] = (_module.__default.fm13(p1, (_this).f16, _1060_v1, globalState)).minus(new BigNumber(-352));
              let _1137_v58;
              let _init28 = ((_1138_v16, _1139_v17) => function (_1140_i4) {
                return _dafny.Seq.of(_1138_v16, _dafny.Seq.of((_1139_v17)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_1139_v17).length))]));
              })(_1083_v16, _1084_v17);
              let _nw170 = Array((new BigNumber(15)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw170.length); _i0_28++) {
                _nw170[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _1137_v58 = _nw170;
              _1137_v58 = _1137_v58;
            } else if (_source19.is_DC8) {
              let _1141___mcc_h6 = (_source19).cf16;
              let _1142___mcc_h7 = (_source19).cf17;
              let _1143___mcc_h8 = (_source19).cf18;
              let _1144_cf18 = _1143___mcc_h8;
              let _1145_cf17 = _1142___mcc_h7;
              let _1146_cf16 = _1141___mcc_h6;
              let _1147_v59;
              _1147_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1144_cf18,new BigNumber((_1120_v49).cardinality()));
              let _1148_v60;
              _1148_v60 = _dafny.MultiSet.fromElements(_1062_v3, (((_1147_v59).contains(_1144_cf18)) ? ((_1147_v59).get(_1144_cf18)) : ((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])), (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))]);
              let _index203 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              (_1117_v46)[_index203] = new BigNumber((_1148_v60).cardinality());
              let _1149_v61;
              _1149_v61 = _dafny.MultiSet.fromElements(_1061_v2, _dafny.Seq.Concat(_1061_v2, _dafny.Seq.UnicodeFromString("hpqdcyt")), (_1116_v45)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-673)), new BigNumber((_1116_v45).length))]);
              _1149_v61 = ((_dafny.MultiSet.FromArray(_1116_v45)).Intersect(_module.__default.fm43(globalState))).Intersect(_1149_v61);
              let _1150_v62;
              _1150_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v3,new BigNumber(767));
              let _1151_v63;
              _1151_v63 = _module.D1.create_DC2(_1150_v62);
              let _1152_v64;
              _1152_v64 = _dafny.Seq.of(_1151_v63, _1151_v63);
              _1152_v64 = _module.__default.fm44(_1062_v3, false, new BigNumber(704), !_dafny.areEqual(_1061_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), ((_1153_v1) => function (_1154_i5) {
                return _1153_v1;
              })(_1060_v1))), globalState);
              let _index204 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              (_1117_v46)[_index204] = (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))];
            } else {
              let _1155___mcc_h9 = (_source19).cf11;
              let _1156_cf11 = _1155___mcc_h9;
              let _index205 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              (_1117_v46)[_index205] = (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))];
              let _1157_v65;
              _1157_v65 = _dafny.Seq.of(_1083_v16, _1083_v16);
              let _1158_v66;
              _1158_v66 = _module.D9.create_DC23(_dafny.MultiSet.FromArray((_1157_v65)[_module.__default.safeIndex(new BigNumber(-343), new BigNumber((_1157_v65).length))]));
              let _1159_v67;
              _1159_v67 = _dafny.MultiSet.fromElements((_this).f16);
              let _1160_v68;
              let _nw171 = new _module.C3();
              _nw171.__ctor(!((_1159_v67).IsProperSubsetOf((_1158_v66).dtor_cf35)), _this.f13);
              _1160_v68 = _nw171;
              let _1161_v69;
              _1161_v69 = _module.D1.create_DC5(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("caoyl"), _module.__default.safeIndex(new BigNumber(435), new BigNumber((_dafny.Seq.UnicodeFromString("caoyl")).length)), _1060_v1), _1084_v17, _1062_v3);
              let _index206 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length));
              (_1117_v46)[_index206] = ((_1161_v69).dtor_cf10).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_1156_cf11)[_module.__default.safeIndex(new BigNumber((_1061_v2).length), new BigNumber((_1156_cf11).length))])));
              let _1162_v70;
              _1162_v70 = _module.D5.create_DC14(p0);
              let _1163_v71;
              _1163_v71 = _module.D5.create_DC15(_1162_v70);
              let _1164_v72;
              _1164_v72 = _dafny.MultiSet.fromElements(_1062_v3, (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))]);
              let _1165_v73;
              _1165_v73 = _dafny.MultiSet.fromElements(_1164_v72, _1164_v72);
              let _index207 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_1084_v17).length));
              (_1084_v17)[_index207] = (_1165_v73).IsProperSubsetOf(_1165_v73);
              let _1166_v74;
              _1166_v74 = _module.D5.create_DC12(_dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1062_v3));
              let _1167_v75;
              _1167_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1061_v2).length),_1084_v17);
              let _1168_v76;
              _1168_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(476),(((_1167_v75).contains((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])) ? ((_1167_v75).get((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])) : (_1084_v17)));
              let _1169_v77;
              _1169_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1084_v17,_1084_v17);
              let _1170_v78;
              let _nw172 = Array((new BigNumber(21)).toNumber());
              _nw172[(_dafny.ZERO).toNumber()] = _1084_v17;
              _nw172[(_dafny.ONE).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(2)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(3)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(4)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(5)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(6)).toNumber()] = (_1161_v69).dtor_cf9;
              _nw172[(new BigNumber(7)).toNumber()] = ((_this.f13) ? (_1084_v17) : (_1084_v17));
              _nw172[(new BigNumber(8)).toNumber()] = (((_1168_v76).contains((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])) ? ((_1168_v76).get((_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))])) : (_1084_v17));
              _nw172[(new BigNumber(9)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(10)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(11)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(12)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(13)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(14)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(15)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(16)).toNumber()] = (((_1169_v77).contains(_1084_v17)) ? ((_1169_v77).get(_1084_v17)) : (_1084_v17));
              _nw172[(new BigNumber(17)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(18)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(19)).toNumber()] = _1084_v17;
              _nw172[(new BigNumber(20)).toNumber()] = _1084_v17;
              _1170_v78 = _nw172;
              let _index208 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1170_v78).length));
              (_1170_v78)[_index208] = _1084_v17;
              let _1171_v79;
              let _nw173 = Array((new BigNumber(2)).toNumber());
              _nw173[(_dafny.ZERO).toNumber()] = _module.__default.fm32(_1061_v2, (_this).f16, _1062_v3, (_1117_v46)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((_1117_v46).length))], globalState);
              _nw173[(_dafny.ONE).toNumber()] = _1061_v2;
              _1171_v79 = _nw173;
              let _index209 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1171_v79).length));
              (_1171_v79)[_index209] = _1061_v2;
              let _1172_v80;
              _1172_v80 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_1061_v2, _1061_v2), _module.__default.safeIndex(new BigNumber(251), new BigNumber((_dafny.Seq.Concat(_1061_v2, _1061_v2)).length)), _1060_v1),!(_dafny.Seq.IsProperPrefixOf(_1061_v2, _1061_v2)));
              let _index210 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_1084_v17).length));
              let _index211 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1170_v78).length));
              let _index212 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1171_v79).length));
              let _rhs177 = _1163_v71;
              let _rhs178 = (((_1172_v80).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), ((_1175_v1) => function (_1176_i6) {
                return _1175_v1;
              })(_1060_v1)))) ? ((_1172_v80).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), ((_1173_v1) => function (_1174_i6) {
                return _1173_v1;
              })(_1060_v1)))) : ((_this).f16));
              let _rhs179 = _1166_v74;
              let _rhs180 = _1084_v17;
              let _rhs181 = _1061_v2;
              let _lhs105 = _1084_v17;
              let _lhs106 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_1084_v17).length));
              let _lhs107 = _1170_v78;
              let _lhs108 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1170_v78).length));
              let _lhs109 = _1171_v79;
              let _lhs110 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1171_v79).length));
              _1163_v71 = _rhs177;
              _lhs105[_lhs106] = _rhs178;
              _1166_v74 = _rhs179;
              _lhs107[_lhs108] = _rhs180;
              _lhs109[_lhs110] = _rhs181;
            }
          }
        }
      }
      let _1177_v81;
      _1177_v81 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
      let _1178_v82;
      _1178_v82 = _dafny.Seq.update(_1083_v16, _module.__default.safeIndex(new BigNumber((_1177_v81).length), new BigNumber((_1083_v16).length)), _this.f13);
      let _1179_v83;
      let _nw174 = Array((new BigNumber(15)).toNumber());
      _nw174[(_dafny.ZERO).toNumber()] = _1178_v82;
      _nw174[(_dafny.ONE).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(2)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(3)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(4)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(5)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(6)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(7)).toNumber()] = _module.__default.fm45(p0, new BigNumber(67), globalState);
      _nw174[(new BigNumber(8)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(9)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(10)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(11)).toNumber()] = _module.__default.fm45(_this.f13, _1062_v3, globalState);
      _nw174[(new BigNumber(12)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(13)).toNumber()] = _1178_v82;
      _nw174[(new BigNumber(14)).toNumber()] = _1178_v82;
      _1179_v83 = _nw174;
      let _index213 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_1179_v83).length));
      (_1179_v83)[_index213] = _1178_v82;
      let _index214 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_1179_v83).length));
      (_1179_v83)[_index214] = _1178_v82;
      let _1180_v84;
      _1180_v84 = _dafny.Set.fromElements(_1062_v3, _1062_v3);
      let _1181_v85;
      _1181_v85 = _dafny.Set.fromElements((_this).f16);
      let _1182_v86;
      _1182_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_1181_v85);
      let _1183_v87;
      _1183_v87 = _dafny.Seq.of(_1181_v85, _1181_v85);
      let _1184_v88;
      _1184_v88 = _dafny.MultiSet.fromElements(new BigNumber((_1061_v2).length));
      let _1185_v89;
      _1185_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1184_v88,p1);
      let _1186_v90;
      let _nw175 = Array((new BigNumber(23)).toNumber());
      _nw175[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(p1, (_1083_v16)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1083_v16).length))]);
      _nw175[(_dafny.ONE).toNumber()] = (((((_1177_v81).contains(p0)) ? ((_1177_v81).get(p0)) : (_this.f13))) ? (_module.__default.fm25(_1180_v84, _1177_v81, _1062_v3, _this.f13, globalState)) : (_dafny.Set.fromElements(p1)));
      _nw175[(new BigNumber(2)).toNumber()] = (((_1182_v86).contains(!(_this.f13))) ? ((_1182_v86).get(!(_this.f13))) : (_dafny.Set.fromElements(p0)));
      _nw175[(new BigNumber(3)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(4)).toNumber()] = ((true) ? (_dafny.Set.fromElements((_this).f16, p0)) : (_1181_v85));
      _nw175[(new BigNumber(5)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(6)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(7)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(8)).toNumber()] = (_1183_v87)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1183_v87).length))];
      _nw175[(new BigNumber(9)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(10)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(11)).toNumber()] = (((((_1185_v89).contains(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_1188_i7) {
        return new BigNumber(915);
      })))) ? ((_1185_v89).get(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_1187_i7) {
        return new BigNumber(915);
      })))) : (p0))) ? (_1181_v85) : (_1181_v85));
      _nw175[(new BigNumber(12)).toNumber()] = (_1181_v85).Intersect(_1181_v85);
      _nw175[(new BigNumber(13)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(14)).toNumber()] = ((p1) ? (_dafny.Set.fromElements(p1, (_this).f16)) : (_1181_v85));
      _nw175[(new BigNumber(15)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(16)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(17)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(18)).toNumber()] = (_1181_v85).Difference(_1181_v85);
      _nw175[(new BigNumber(19)).toNumber()] = ((_this.f13) ? (_1181_v85) : (_dafny.Set.fromElements((_this).f16)));
      _nw175[(new BigNumber(20)).toNumber()] = _1181_v85;
      _nw175[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(true);
      _nw175[(new BigNumber(22)).toNumber()] = _1181_v85;
      _1186_v90 = _nw175;
      let _index215 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1186_v90).length));
      (_1186_v90)[_index215] = (_1181_v85).Intersect(_1181_v85);
      let _index216 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1186_v90).length));
      let _rhs182 = _1084_v17;
      let _rhs183 = ((_1181_v85).Union(_dafny.Set.fromElements((_1083_v16)[_module.__default.safeIndex(_1062_v3, new BigNumber((_1083_v16).length))], _this.f13, p0, (_this).f16))).Union(_1181_v85);
      let _lhs111 = _1186_v90;
      let _lhs112 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1186_v90).length));
      _1084_v17 = _rhs182;
      _lhs111[_lhs112] = _rhs183;
      r0 = _dafny.Seq.IsPrefixOf(_1061_v2, _1061_v2);
      r1 = (_1181_v85).IsProperSubsetOf(((_1186_v90)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1186_v90).length))]).Intersect(_1181_v85));
      r2 = (new BigNumber((_1177_v81).length)).plus(_1062_v3);
      return [r0, r1, r2];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f13 = false;
      this._f14 = [];
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    set f13(value) {
      let _this = this;
      _this._f13 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f13, f14) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((_this.f13) ? (new BigNumber(790)) : (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f13)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13)).length))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-516), new BigNumber(198), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("flisrw")).length)), new BigNumber((_dafny.Set.fromElements(_this.f13)).length))).length))).length)))).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(757),new BigNumber(-211))).length));
    };
    fm5(globalState) {
      let _this = this;
      return _module.D1.create_DC3(_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_this.f13)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(346)), function (_1189_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length)));
    };
    fm6(globalState) {
      let _this = this;
      return !((_module.D0.create_DC1(_this.f13, new BigNumber((_dafny.Seq.UnicodeFromString("box")).length))).dtor_cf1);
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber(457), new BigNumber(-669))).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("nuqc")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_module.D0.create_DC1(false, new BigNumber(114)))).length)));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("rq"), new _dafny.CodePoint('a'.codePointAt(0)));
    };
    fm9(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_dafny.ZERO).minus(((_this.f13) ? (new BigNumber(-694)) : (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f13, !(_this.f13), _this.f13))).cardinality())))), (new BigNumber(-159)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(!(_this.f13), _this.f13)).cardinality())));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let _1190_v0;
      let _nw176 = Array((_dafny.ONE).toNumber());
      _nw176[(_dafny.ZERO).toNumber()] = !(!(_this.f13));
      _1190_v0 = _nw176;
      let _1191_v1;
      _1191_v1 = _module.D1.create_DC4(_1190_v0, p0, new _dafny.CodePoint('i'.codePointAt(0)));
      let _source20 = _1191_v1;
      if (_source20.is_DC3) {
        let _1192___mcc_h0 = (_source20).cf4;
        let _1193_cf4 = _1192___mcc_h0;
        let _1194_v2;
        _1194_v2 = _dafny.MultiSet.fromElements(_1193_cf4);
        _1193_cf4 = (_this).fm7((_this).fm9(p0, globalState), _module.__default.safeDivisionInt(p0, new BigNumber(((_1194_v2).update(p0, _module.__default.abs(p0))).cardinality())), _this.f13, globalState);
        if (!(true)) {
          let _1195_v3;
          _1195_v3 = _dafny.Seq.UnicodeFromString("gljnndblh");
          let _1196_v4;
          let _out31;
          _out31 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_1195_v3,_this.f13), globalState);
          _1196_v4 = _out31;
          let _1197_v5;
          let _nw177 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
          _1197_v5 = _nw177;
          _1197_v5 = _1197_v5;
          let _1198_v6;
          let _nw178 = new _module.C0();
          _nw178.__ctor((_this).fm7(_1193_cf4, (_dafny.ZERO).minus(_1196_v4), _this.f13, globalState));
          _1198_v6 = _nw178;
          let _1199_v7;
          let _1200_v8;
          let _out32;
          let _out33;
          let _outcollector6 = (_this).m3(_1193_cf4, (_1198_v6).f15, (_1198_v6).f15, globalState);
          _out32 = _outcollector6[0];
          _out33 = _outcollector6[1];
          _1199_v7 = _out32;
          _1200_v8 = _out33;
          _1195_v3 = _1195_v3;
        } else {
          let _rhs184 = _this.f13;
          let _rhs185 = _1193_cf4;
          let _lhs113 = _this;
          _lhs113.f13 = _rhs184;
          _1193_cf4 = _rhs185;
          let _1201_v9;
          _1201_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,(_1194_v2).update(p0, _module.__default.abs(p0)));
          _1201_v9 = (_1201_v9).update(_this.f13, (_1194_v2).Difference(_1194_v2));
          _1193_cf4 = p0;
          let _index217 = _module.__default.safeIndex(new BigNumber(371), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index217] = p0;
          let _index218 = _module.__default.safeIndex(new BigNumber(371), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index218] = _1193_cf4;
          let _1202_v10;
          _1202_v10 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f14)[_module.__default.safeIndex(new BigNumber(371), new BigNumber(((_this).f14).length))],_this.f13);
          _1193_cf4 = new BigNumber((_1202_v10).length);
        }
        let _1203_v11;
        _1203_v11 = _dafny.Seq.of((_this).f14);
        let _1204_v12;
        let _nw179 = new _module.C0();
        _nw179.__ctor((_dafny.ZERO).minus(new BigNumber((_1203_v11).length)));
        _1204_v12 = _nw179;
        let _1205_v13;
        _1205_v13 = _dafny.Seq.UnicodeFromString("dqmf");
        let _1206_v14;
        _1206_v14 = new _dafny.CodePoint('w'.codePointAt(0));
        _1205_v13 = _dafny.Seq.update(_dafny.Seq.Concat(_1205_v13, _dafny.Seq.UnicodeFromString("abddeo")), _module.__default.safeIndex((p0).multipliedBy(new BigNumber((_1205_v13).length)), new BigNumber((_dafny.Seq.Concat(_1205_v13, _dafny.Seq.UnicodeFromString("abddeo"))).length)), _1206_v14);
      } else if (_source20.is_DC4) {
        let _1207___mcc_h1 = (_source20).cf5;
        let _1208___mcc_h2 = (_source20).cf6;
        let _1209___mcc_h3 = (_source20).cf7;
        let _1210_cf7 = _1209___mcc_h3;
        let _1211_cf6 = _1208___mcc_h2;
        let _1212_cf5 = _1207___mcc_h1;
        let _1213_v15;
        _1213_v15 = _dafny.MultiSet.fromElements(_this.f13, _this.f13);
        _1211_cf6 = new BigNumber((_1213_v15).cardinality());
        (_this).f13 = _this.f13;
        let _1214_v16;
        _1214_v16 = _dafny.Seq.UnicodeFromString("buj");
        _1214_v16 = _1214_v16;
        let _1215_v17;
        _1215_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1211_cf6,(_this).fm9((_this).fm9(p0, globalState), globalState));
        let _1216_v18;
        _1216_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(685),_1215_v17);
        let _1217_v19;
        _1217_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
        _1216_v18 = (_1216_v18).update((_dafny.ZERO).minus((_this).fm9(new BigNumber((_1217_v19).length), globalState)), (_1215_v17).update(p0, new BigNumber((_1213_v15).cardinality())));
      } else if (_source20.is_DC5) {
        let _1218___mcc_h4 = (_source20).cf8;
        let _1219___mcc_h5 = (_source20).cf9;
        let _1220___mcc_h6 = (_source20).cf10;
        let _1221_cf10 = _1220___mcc_h6;
        let _1222_cf9 = _1219___mcc_h5;
        let _1223_cf8 = _1218___mcc_h4;
        let _1224_v20;
        let _init29 = ((_1225_p0) => function (_1226_i0) {
          return _dafny.MultiSet.FromArray((_module.D2.create_DC6(_dafny.Seq.of(_1225_p0))).dtor_cf11);
        })(p0);
        let _nw180 = Array((new BigNumber(25)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw180.length); _i0_29++) {
          _nw180[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1224_v20 = _nw180;
        _1224_v20 = _1224_v20;
        let _1227_v21;
        _1227_v21 = _module.D0.create_DC0(_this.f13);
        _1227_v21 = _1227_v21;
        let _1228_v22;
        _1228_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1191_v1);
        let _1229_v23;
        _1229_v23 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new _dafny.CodePoint('x'.codePointAt(0)));
        _1228_v22 = (_1228_v22).update(new BigNumber((_1229_v23).length), _1191_v1);
        let _1230_v24;
        _1230_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1223_cf8,_module.__default.fm11(_this.f13, p0, _this.f13, globalState));
        let _1231_v25;
        _1231_v25 = new _dafny.CodePoint('b'.codePointAt(0));
        _1230_v24 = (_1230_v24).update(_1223_cf8, _1231_v25);
      } else {
        let _1232___mcc_h7 = (_source20).cf3;
        let _1233_cf3 = _1232___mcc_h7;
        let _1234_v26;
        _1234_v26 = new BigNumber(283);
        _1234_v26 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1234_v26));
        _1234_v26 = p0;
        _1234_v26 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1234_v26,((!(_this.f13)) ? (_1234_v26) : (_1234_v26)))).length);
        (_this).f13 = _this.f13;
      }
      let _1235_v27;
      _1235_v27 = new BigNumber(-407);
      let _1236_v28;
      _1236_v28 = _dafny.Seq.UnicodeFromString("sffwtcir");
      _1235_v27 = new BigNumber((_dafny.Seq.Concat(_1236_v28, _1236_v28)).length);
      let _1237_v29;
      let _nw181 = Array((new BigNumber(18)).toNumber());
      _1237_v29 = _nw181;
      let _1238_v30;
      _1238_v30 = _dafny.Set.fromElements(_this.f13);
      let _1239_v31;
      _1239_v31 = _dafny.Set.fromElements(_1238_v30);
      let _1240_v32;
      let _nw182 = new _module.C0();
      _nw182.__ctor(new BigNumber((_1239_v31).length));
      _1240_v32 = _nw182;
      let _index219 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1237_v29).length));
      (_1237_v29)[_index219] = _1240_v32;
      let _index220 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1237_v29).length));
      (_1237_v29)[_index220] = _1240_v32;
      (_this).f13 = true;
      let _1241_v33;
      let _nw183 = new _module.C0();
      _nw183.__ctor(p0);
      _1241_v33 = _nw183;
      let _1242_v34;
      _1242_v34 = _dafny.Set.fromElements(new BigNumber(549));
      let _1243_v35;
      _1243_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-689)),_1242_v34);
      let _1244_v36;
      _1244_v36 = _dafny.MultiSet.fromElements(false, _this.f13);
      let _1245_v37;
      _1245_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1244_v36,(_1241_v33).f15);
      let _1246_v38;
      _1246_v38 = _dafny.Seq.of(p0);
      let _1247_v39;
      _1247_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1246_v38)[_module.__default.safeIndex((_this).fm9(_1235_v27, globalState), new BigNumber((_1246_v38).length))],(_this).f14);
      let _rhs186 = (new BigNumber((((((_1243_v35).contains((_1240_v32).f15)) ? ((_1243_v35).get((_1240_v32).f15)) : (_1242_v34))).Difference(_1242_v34)).length)).isLessThanOrEqualTo(new BigNumber((_1245_v37).length));
      let _rhs187 = _1241_v33;
      let _rhs188 = (p0).multipliedBy(new BigNumber((_1247_v39).length));
      let _lhs114 = _this;
      _lhs114.f13 = _rhs186;
      _1240_v32 = _rhs187;
      _1235_v27 = _rhs188;
      r0 = (_this).f14;
      return r0;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _1248_i0;
      _1248_i0 = _dafny.ZERO;
      L11: {
        while (_this.f13) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1248_i0)) {
              break L11;
            }
            _1248_i0 = (_1248_i0).plus(_dafny.ONE);
            let _1249_v0;
            _1249_v0 = new BigNumber(21);
            _1249_v0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p2, p2), ((_this.f13) ? (p0) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), function (_1250_i1) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length))));
            let _1251_v1;
            let _init30 = ((_1252_p1) => function (_1253_i2) {
              return _dafny.Seq.Concat(_1252_p1, _1252_p1);
            })(p1);
            let _nw184 = Array((new BigNumber(20)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw184.length); _i0_30++) {
              _nw184[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _1251_v1 = _nw184;
            let _index221 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1251_v1).length));
            (_1251_v1)[_index221] = p1;
            let _index222 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1251_v1).length));
            (_1251_v1)[_index222] = p1;
            let _1254_v2;
            _1254_v2 = _dafny.MultiSet.fromElements(p2);
            let _1255_v3;
            _1255_v3 = _dafny.Set.fromElements(_1254_v2);
            let _1256_v4;
            _1256_v4 = new _dafny.CodePoint('t'.codePointAt(0));
            let _1257_v5;
            _1257_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1255_v3,_dafny.Seq.Concat(p1, _dafny.Seq.update(p1, _module.__default.safeIndex(_1249_v0, new BigNumber((p1).length)), _1256_v4)));
            _1257_v5 = (_1257_v5).update(_1255_v3, _dafny.Seq.update((_1251_v1)[_module.__default.safeIndex(new BigNumber(152), new BigNumber((_1251_v1).length))], _module.__default.safeIndex(p2, new BigNumber(((_1251_v1)[_module.__default.safeIndex(new BigNumber(152), new BigNumber((_1251_v1).length))]).length)), _1256_v4));
            let _1258_v6;
            let _nw185 = Array((new BigNumber(12)).toNumber());
            _1258_v6 = _nw185;
            let _1259_v7;
            let _nw186 = new _module.C3();
            _nw186.__ctor(_this.f13, _this.f13);
            _1259_v7 = _nw186;
            let _index223 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1258_v6).length));
            (_1258_v6)[_index223] = _1259_v7;
            let _index224 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1258_v6).length));
            let _nw187 = new _module.C3();
            _nw187.__ctor(_this.f13, !(_this.f13));
            (_1258_v6)[_index224] = _nw187;
          }
        }
      }
      let _1260_v8;
      let _nw188 = new _module.C3();
      _nw188.__ctor(((_this.f13) ? (false) : (_this.f13)), _this.f13);
      _1260_v8 = _nw188;
      if (_this.f13) {
        let _1261_v9;
        _1261_v9 = _dafny.Seq.of(!(_this.f13));
        let _1262_v10;
        _1262_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1261_v9).length));
        let _1263_v11;
        _1263_v11 = _module.D0.create_DC1(_this.f13, (((_1262_v10).contains(p2)) ? ((_1262_v10).get(p2)) : (p0)));
        let _source21 = _1263_v11;
        if (_source21.is_DC1) {
          let _1264___mcc_h0 = (_source21).cf1;
          let _1265___mcc_h1 = (_source21).cf2;
          let _1266_cf2 = _1265___mcc_h1;
          let _1267_cf1 = _1264___mcc_h0;
          (_this).f13 = _module.__default.fm0((_1262_v10).Merge(_1262_v10), (new BigNumber((p1).length)).multipliedBy(p0), globalState);
          _1266_cf2 = new BigNumber(207);
          _1266_cf2 = p0;
          let _1268_v12;
          _1268_v12 = _dafny.Set.fromElements(p2, p0, p0);
          let _1269_v13;
          _1269_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1261_v9).length),(_1260_v8).f19);
          let _1270_v15;
          let _nw189 = Array((new BigNumber(5)).toNumber());
          _nw189[(_dafny.ZERO).toNumber()] = _1268_v12;
          _nw189[(_dafny.ONE).toNumber()] = _1268_v12;
          _nw189[(new BigNumber(2)).toNumber()] = (_1268_v12).Intersect(_1268_v12);
          _nw189[(new BigNumber(3)).toNumber()] = _1268_v12;
          _nw189[(new BigNumber(4)).toNumber()] = function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of (_1269_v13).Keys.Elements) {
              let _1271_v14 = _compr_42;
              if ((_1269_v13).contains(_1271_v14)) {
                _coll42.add((_1271_v14).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, false))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), function (_1272_i3) {
                  return new _dafny.CodePoint('o'.codePointAt(0));
                })).length))).length)));
              }
            }
            return _coll42;
          }();
          _1270_v15 = _nw189;
          let _1273_v16;
          _1273_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1260_v8);
          let _index225 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1270_v15).length));
          (_1270_v15)[_index225] = _dafny.Set.fromElements(new BigNumber((_1273_v16).length));
          let _1274_v17;
          _1274_v17 = _dafny.Seq.UnicodeFromString("ucptmuvvm");
          let _1275_v18;
          _1275_v18 = _dafny.Seq.of(p2);
          let _index226 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1270_v15).length));
          let _rhs189 = ((_1268_v12).Union(_1268_v12)).Difference((_dafny.Set.fromElements(_1266_cf2, _1266_cf2, new BigNumber((_1275_v18).length))).Intersect(_1268_v12));
          let _rhs190 = _this.f13;
          let _rhs191 = (_1260_v8).f19;
          let _rhs192 = new BigNumber((_dafny.Set.fromElements((p0).plus(_1266_cf2), (_dafny.ZERO).minus((_dafny.ZERO).minus(p2)), p0)).length);
          let _rhs193 = p1;
          let _lhs115 = _1270_v15;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1270_v15).length));
          _lhs115[_lhs116] = _rhs189;
          _1267_cf1 = _rhs190;
          _1267_cf1 = _rhs191;
          _1266_cf2 = _rhs192;
          _1274_v17 = _rhs193;
        } else {
          let _1276___mcc_h2 = (_source21).cf0;
          let _1277_cf0 = _1276___mcc_h2;
          let _1278_v19;
          _1278_v19 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1279_v20;
          _1279_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1260_v8).f19,(_1260_v8).f19);
          let _1280_v21;
          _1280_v21 = _dafny.MultiSet.fromElements((((_1279_v20).contains(_1277_cf0)) ? ((_1279_v20).get(_1277_cf0)) : ((_1260_v8).f19)), _this.f13);
          let _1281_v22;
          _1281_v22 = _module.D2.create_DC7(_dafny.Seq.update(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), _1278_v19, _1278_v19), _module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), _1278_v19, _1278_v19)).length)), _1278_v19), new BigNumber((_1280_v21).cardinality()), p2, _1277_cf0);
          let _1282_v23;
          _1282_v23 = _dafny.Map.Empty.slice().updateUnsafe(!((_1281_v22).dtor_cf15),p2);
          _1282_v23 = (_1282_v23).update(!((_1260_v8).f19), p2);
          let _1283_v24;
          _1283_v24 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1262_v10);
          let _1284_v25;
          let _nw190 = new _module.C4();
          _nw190.__ctor((_this).fm8(_module.D1.create_DC2((((_1283_v24).contains(p2)) ? ((_1283_v24).get(p2)) : (_module.__default.fm18((_1260_v8).f19, (_1260_v8).f19, new BigNumber((_1282_v23).length), (_this).fm7(p0, p2, _1277_cf0, globalState), globalState)))), p2, globalState), !((_1260_v8).f19));
          _1284_v25 = _nw190;
          let _1285_v26;
          _1285_v26 = _module.D7.create_DC20(new BigNumber(606), _1278_v19);
          let _1286_v27;
          _1286_v27 = _module.D7.create_DC21(_1285_v26);
          let _1287_v28;
          _1287_v28 = _module.D7.create_DC21(_1285_v26);
          let _1288_v29;
          _1288_v29 = _module.D7.create_DC21(_1286_v27);
          _1288_v29 = _1288_v29;
          _1279_v20 = (_1279_v20).Merge((_1279_v20).Merge(_1279_v20));
        }
        let _1289_v30;
        _1289_v30 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1290_v31;
        _1290_v31 = _dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC11((_1260_v8).f19)).dtor_cf21,_this.f13);
        let _1291_v32;
        _1291_v32 = _module.D2.create_DC8(_this.f13, false, _this.f13);
        _1289_v30 = _module.__default.fm11((((_1290_v31).contains((_1260_v8).f19)) ? ((_1290_v31).get((_1260_v8).f19)) : ((_1291_v32).dtor_cf17)), p2, _module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-513)), p2, globalState), globalState);
        let _1292_v33;
        let _nw191 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1292_v33 = _nw191;
        let _index227 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        (_1292_v33)[_index227] = (p0).isLessThanOrEqualTo(p2);
        let _1293_v34;
        let _nw192 = new _module.C4();
        _nw192.__ctor(_this.f13, (_1260_v8).f19);
        _1293_v34 = _nw192;
        let _1294_v35;
        _1294_v35 = _dafny.Set.fromElements((_1260_v8).f19);
        let _1295_v36;
        _1295_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1293_v34,_1294_v35);
        let _index228 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        let _rhs194 = (_1260_v8).f19;
        let _rhs195 = (p0).isLessThanOrEqualTo((_this).fm9(new BigNumber(((((_1295_v36).contains(_1293_v34)) ? ((_1295_v36).get(_1293_v34)) : (_dafny.Set.fromElements(_this.f13)))).length), globalState));
        let _lhs117 = _this;
        let _lhs118 = _1292_v33;
        let _lhs119 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        _lhs117.f13 = _rhs194;
        _lhs118[_lhs119] = _rhs195;
        let _1296_v37;
        _1296_v37 = _1261_v9;
        let _source22 = _1296_v37;
        let _1297___mcc_h3 = _source22;
        let _1298_cf34 = _1297___mcc_h3;
        let _index229 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        (_1292_v33)[_index229] = (_1260_v8).f19;
        let _index230 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index230] = p2;
        let _index231 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index231] = p2;
        let _1299_v38;
        _1299_v38 = _module.D5.create_DC14((_1292_v33)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length))]);
        let _1300_v39;
        _1300_v39 = _dafny.Set.fromElements(_module.D5.create_DC14((_1293_v34).f16), _1299_v38, _1299_v38, _1299_v38);
        let _1301_v40;
        let _nw193 = Array((new BigNumber(9)).toNumber());
        _nw193[(_dafny.ZERO).toNumber()] = _1292_v33;
        _nw193[(_dafny.ONE).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(2)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(3)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(4)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(5)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(6)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(7)).toNumber()] = _1292_v33;
        _nw193[(new BigNumber(8)).toNumber()] = _1292_v33;
        _1301_v40 = _nw193;
        let _1302_v41;
        _1302_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1300_v39,_1301_v40);
        _1302_v41 = (_1302_v41).update(_1300_v39, _1301_v40);
        let _1303_v42;
        let _init31 = ((_1304_v34) => function (_1305_i4) {
          return _module.D9.create_DC23(_dafny.MultiSet.fromElements((_1304_v34).f16));
        })(_1293_v34);
        let _nw194 = Array((new BigNumber(25)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw194.length); _i0_31++) {
          _nw194[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1303_v42 = _nw194;
        let _1306_v43;
        _1306_v43 = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm7(((_this).f14)[_module.__default.safeIndex(new BigNumber(980), new BigNumber(((_this).f14).length))], new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0))).cardinality()), (_1293_v34).f16, globalState)));
        let _1307_v44;
        _1307_v44 = _dafny.MultiSet.fromElements((_this).fm8(_1306_v43, p0, globalState), (_1293_v34).f16);
        let _1308_v45;
        _1308_v45 = _module.D9.create_DC23(_1307_v44);
        let _index232 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1303_v42).length));
        (_1303_v42)[_index232] = _1308_v45;
        let _index233 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        let _index234 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_this).f14).length));
        let _index235 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1303_v42).length));
        let _rhs196 = !(false);
        let _rhs197 = (((_1307_v44).contains(!(true))) ? ((_1307_v44).get(!(true))) : (p2));
        let _rhs198 = _1308_v45;
        let _rhs199 = _1289_v30;
        let _lhs120 = _1292_v33;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v33).length));
        let _lhs122 = (_this).f14;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(980), new BigNumber(((_this).f14).length));
        let _lhs124 = _1303_v42;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1303_v42).length));
        _lhs120[_lhs121] = _rhs196;
        _lhs122[_lhs123] = _rhs197;
        _lhs124[_lhs125] = _rhs198;
        _1289_v30 = _rhs199;
        _1261_v9 = _1261_v9;
      } else {
        let _1309_v46;
        _1309_v46 = _dafny.Seq.of(p0, new BigNumber(627), p0, new BigNumber((p1).length));
        let _1310_v47;
        let _nw195 = new _module.C3();
        _nw195.__ctor((((_1260_v8).f19) ? (true) : (true)), !_dafny.areEqual(_1309_v46, _1309_v46));
        _1310_v47 = _nw195;
        let _1311_v48;
        _1311_v48 = new BigNumber(58);
        let _1312_v49;
        _1312_v49 = new _dafny.CodePoint('r'.codePointAt(0));
        _1311_v48 = _module.__default.fm13((_1260_v8).f19, _this.f13, _1312_v49, globalState);
        _1311_v48 = p0;
        let _1313_v50;
        let _nw196 = new _module.C3();
        _nw196.__ctor(!(_dafny.Set.fromElements(false)).contains(_this.f13), !(p0).isEqualTo((_dafny.ZERO).minus(p0)));
        _1313_v50 = _nw196;
        let _1314_v51;
        let _nw197 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
        _1314_v51 = _nw197;
        let _1315_v52;
        let _nw198 = new _module.C0();
        _nw198.__ctor(p2);
        _1315_v52 = _nw198;
        let _1316_v53;
        _1316_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1315_v52,(_1310_v47).f19);
        let _index236 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1314_v51).length));
        (_1314_v51)[_index236] = _1316_v53;
        let _index237 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1314_v51).length));
        (_1314_v51)[_index237] = ((_1316_v53).update(_1315_v52, (_1260_v8).f19)).Merge(_1316_v53);
      }
      let _1317_v54;
      _1317_v54 = _dafny.Seq.UnicodeFromString("jvtqh");
      _1317_v54 = p1;
      let _index238 = _module.__default.safeIndex(new BigNumber(400), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index238] = p2;
      let _1318_v55;
      _1318_v55 = _dafny.MultiSet.fromElements(p0, p0);
      let _1319_v56;
      _1319_v56 = _dafny.Seq.of(_dafny.Seq.of(p0, new BigNumber((_1318_v55).cardinality()), p0));
      let _index239 = _module.__default.safeIndex(new BigNumber(400), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index239] = new BigNumber((_1319_v56).length);
      let _1320_v57;
      _1320_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1321_v58;
      _1321_v58 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f14)[_module.__default.safeIndex(new BigNumber(400), new BigNumber(((_this).f14).length))],_module.__default.fm0(_1320_v57, new BigNumber((_dafny.MultiSet.fromElements(_this.f13)).cardinality()), globalState));
      let _1322_v59;
      _1322_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1321_v58).Merge(_1321_v58),_1317_v54);
      _1322_v59 = (_1322_v59).update(_1321_v58, _dafny.Seq.UnicodeFromString("hbsuyxpfb"));
      let _1323_v60;
      let _init32 = function (_1324_i5) {
        return _this.f13;
      };
      let _nw199 = Array((new BigNumber(20)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw199.length); _i0_32++) {
        _nw199[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1323_v60 = _nw199;
      let _1325_v61;
      _1325_v61 = _dafny.Seq.of(_1323_v60);
      let _1326_v62;
      _1326_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(73),(_1325_v61)[_module.__default.safeIndex(p2, new BigNumber((_1325_v61).length))]);
      let _1327_v63;
      _1327_v63 = _dafny.Seq.of(p2);
      let _1328_v64;
      _1328_v64 = _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("yvhsai"), _1323_v60, p2);
      let _1329_v65;
      _1329_v65 = _dafny.Seq.of((_1326_v62).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1327_v63).length),(_module.D1.create_DC5((_1328_v64).dtor_cf8, _1323_v60, ((_this).f14)[_module.__default.safeIndex(new BigNumber(400), new BigNumber(((_this).f14).length))])).dtor_cf9)));
      r0 = (_1329_v65)[_module.__default.safeIndex(((_this).f14)[_module.__default.safeIndex(new BigNumber(400), new BigNumber(((_this).f14).length))], new BigNumber((_1329_v65).length))];
      let _1330_v66;
      _1330_v66 = _dafny.Seq.of((_1260_v8).f19);
      r1 = !(p0).isEqualTo((new BigNumber((_1330_v66).length)).minus((_this).fm9(new BigNumber((_1327_v63).length), globalState)));
      r2 = (((_1260_v8).f19) ? ((_1260_v8).f19) : (_this.f13));
      r3 = _1330_v66;
      return [r0, r1, r2, r3];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1331_v0;
      _1331_v0 = new _dafny.CodePoint('m'.codePointAt(0));
      let _1332_v1;
      _1332_v1 = _module.D7.create_DC20(p1, _1331_v0);
      r1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, (_1332_v1).dtor_cf31));
      if (!(true)) {
        r0 = _this.f13;
        let _1333_v2;
        _1333_v2 = _dafny.MultiSet.fromElements(new BigNumber(508));
        r1 = new BigNumber((_1333_v2).cardinality());
        let _1334_v3;
        let _nw200 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1334_v3 = _nw200;
        let _1335_v4;
        _1335_v4 = _dafny.Seq.of(_this.f13, _this.f13);
        let _1336_v5;
        _1336_v5 = _dafny.Set.fromElements(new BigNumber((_1335_v4).length), new BigNumber((_1335_v4).length));
        let _index240 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1334_v3).length));
        (_1334_v3)[_index240] = (_1336_v5).IsDisjointFrom(_dafny.Set.fromElements((_this).fm9(new BigNumber(165), globalState), p1));
        let _1337_v6;
        let _nw201 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1337_v6 = _nw201;
        let _1338_v7;
        _1338_v7 = _dafny.Seq.UnicodeFromString("n");
        let _index241 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1337_v6).length));
        (_1337_v6)[_index241] = _dafny.Seq.Concat(_1338_v7, _1338_v7);
        let _1339_v8;
        _1339_v8 = _module.D1.create_DC5(_1338_v7, _1334_v3, p0);
        let _index242 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1334_v3).length));
        let _index243 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1337_v6).length));
        let _rhs200 = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ywfkh"), _1331_v0);
        let _rhs201 = (_1339_v8).dtor_cf8;
        let _lhs126 = _1334_v3;
        let _lhs127 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1334_v3).length));
        let _lhs128 = _1337_v6;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_1337_v6).length));
        _lhs126[_lhs127] = _rhs200;
        _lhs128[_lhs129] = _rhs201;
        _1338_v7 = _1338_v7;
        let _1340_v9;
        let _nw202 = Array((new BigNumber(3)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw202[(_dafny.ONE).toNumber()] = _1331_v0;
        _nw202[(new BigNumber(2)).toNumber()] = _1331_v0;
        _1340_v9 = _nw202;
        let _index244 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1340_v9).length));
        (_1340_v9)[_index244] = _1331_v0;
        let _index245 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1340_v9).length));
        (_1340_v9)[_index245] = new _dafny.CodePoint('y'.codePointAt(0));
      } else {
        r1 = (new BigNumber(-939)).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f13, _this.f13, true)).length));
        let _1341_v10;
        _1341_v10 = _dafny.Seq.UnicodeFromString("garuvy");
        let _1342_v11;
        _1342_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1341_v10,!(_this.f13) || ((_this).fm6(globalState)));
        _1342_v11 = (_1342_v11).update(_dafny.Seq.UnicodeFromString("fg"), !(_this.f13));
        let _1343_v12;
        _1343_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        let _1344_v13;
        _1344_v13 = _dafny.MultiSet.fromElements(_this.f13);
        let _1345_v14;
        _1345_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_this.f13)).length),_1344_v13);
        _1343_v12 = (_1343_v12).update(new BigNumber(-5), !(_1345_v14).contains(p0));
        let _1346_v15;
        _1346_v15 = _module.D0.create_DC0(_this.f13);
        _1346_v15 = _module.D0.create_DC0(_this.f13);
        r1 = _module.__default.safeModuloInt(p0, new BigNumber(898));
      }
      let _1347_v16;
      let _nw203 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _1347_v16 = _nw203;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1347_v16).length))) {
        let _1348_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1348_i0)) && ((_1348_i0).isLessThan(new BigNumber((_1347_v16).length))))) {
          (_1347_v16)[(_1348_i0)] = (_1348_i0).plus(p2);
        }
      }
      let _1349_v17;
      _1349_v17 = _dafny.Seq.UnicodeFromString("isf");
      let _hi11 = new BigNumber((_1349_v17).length);
      for (let _1350_i1 = p0; _1350_i1.isLessThan(_hi11); _1350_i1 = _1350_i1.plus(_dafny.ONE)) {
        let _1351_v18;
        let _init33 = function (_1352_i2) {
          return false;
        };
        let _nw204 = Array((new BigNumber(11)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw204.length); _i0_33++) {
          _nw204[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1351_v18 = _nw204;
        let _index246 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length));
        (_1351_v18)[_index246] = _this.f13;
        let _1353_v19;
        _1353_v19 = _dafny.Seq.of(!(p0).isEqualTo(p1));
        let _1354_v20;
        _1354_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
        let _index247 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length));
        (_1351_v18)[_index247] = (_1353_v19)[_module.__default.safeIndex(new BigNumber((_1354_v20).length), new BigNumber((_1353_v19).length))];
        let _1355_v21;
        _1355_v21 = _dafny.Seq.of(p2);
        let _index248 = _module.__default.safeIndex(new BigNumber(560), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index248] = ((_1355_v21)[_module.__default.safeIndex(new BigNumber((_1349_v17).length), new BigNumber((_1355_v21).length))]).multipliedBy(p1);
        let _index249 = _module.__default.safeIndex(new BigNumber(560), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index249] = (_dafny.ZERO).minus(p2);
        (_this).f13 = (_1351_v18)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length))];
        let _1356_v23;
        let _nw205 = new _module.C4();
        _nw205.__ctor(true, _this.f13);
        _1356_v23 = _nw205;
        let _1357_v24;
        _1357_v24 = _dafny.MultiSet.fromElements((_1356_v23).f16);
        let _1358_v25;
        _1358_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1356_v23,_1357_v24);
        let _1359_v26;
        _1359_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f13,p2)).length),(_dafny.ZERO).minus(new BigNumber((_1358_v25).length)));
        let _1360_v27;
        _1360_v27 = _dafny.Set.fromElements((((_1351_v18)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length))]) ? (function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(983), new BigNumber(402))) {
            let _1361_v22 = _compr_43;
            if (((new BigNumber(983)).isLessThanOrEqualTo(_1361_v22)) && ((_1361_v22).isLessThan(new BigNumber(402)))) {
              _coll43.push([(_1361_v22).minus(p0),p0]);
            }
          }
          return _coll43;
        }()) : ((_1359_v26).update(p0, new BigNumber(361)))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1349_v17).length),new BigNumber(726)), (_1359_v26).update(p0, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_module.__default.fm13((_1351_v18)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length))], (_1351_v18)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length))], _1331_v0, globalState), p0, p2, new BigNumber((_module.__default.fm24(new BigNumber((_1349_v17).length), (_1351_v18)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length))], globalState)).length))).length))), _dafny.Map.Empty.slice().updateUnsafe(p2,p2));
        let _1362_v28;
        _1362_v28 = _module.D0.create_DC1((_1356_v23).f16, new BigNumber(356));
        let _index250 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length));
        let _rhs202 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(560), new BigNumber(((_this).f14).length))];
        let _rhs203 = p0;
        let _rhs204 = (_1356_v23).fm4(((_dafny.ZERO).minus(p1)).isLessThan(new BigNumber(306)), _1362_v28, _1350_i1, _1359_v26, globalState);
        let _rhs205 = _module.__default.fm46(p0, globalState);
        let _lhs130 = _1351_v18;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1351_v18).length));
        r1 = _rhs202;
        r1 = _rhs203;
        _lhs130[_lhs131] = _rhs204;
        _1360_v27 = _rhs205;
      }
      let _1363_i3;
      _1363_i3 = _dafny.ZERO;
      L12: {
        while (_this.f13) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1363_i3)) {
              break L12;
            }
            _1363_i3 = (_1363_i3).plus(_dafny.ONE);
            r1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p2, ((true) ? (new BigNumber(-158)) : (p0))));
            let _1364_v29;
            _1364_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,(true) === (true));
            let _1365_v30;
            _1365_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
            let _1366_v31;
            _1366_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1365_v30).length)),p2);
            _1364_v29 = (_1364_v29).update(((_module.__default.fm0(_1366_v31, p0, globalState)) ? (p1) : (p0)), _this.f13);
            let _1367_v32;
            _1367_v32 = _module.D2.create_DC7(_1349_v17, p1, p0, _this.f13);
            let _1368_v33;
            _1368_v33 = _dafny.Seq.of(_this.f13, _this.f13, _this.f13, _this.f13, (_this).fm4(_this.f13, _module.D0.create_DC1(_this.f13, new BigNumber(-46)), new BigNumber(675), _1366_v31, globalState));
            let _1369_v34;
            _1369_v34 = _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0))));
            let _pat_let_tv24 = p1;
            let _1370_v35;
            let _nw206 = Array((new BigNumber(19)).toNumber());
            _nw206[(_dafny.ZERO).toNumber()] = _1367_v32;
            _nw206[(_dafny.ONE).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(2)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(3)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(4)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(5)).toNumber()] = function (_pat_let28_0) {
              return function (_1371_dt__update__tmp_h0) {
                return function (_pat_let29_0) {
                  return function (_1372_dt__update_hcf13_h0) {
                    return _module.D2.create_DC7((_1371_dt__update__tmp_h0).dtor_cf12, _1372_dt__update_hcf13_h0, (_1371_dt__update__tmp_h0).dtor_cf14, (_1371_dt__update__tmp_h0).dtor_cf15);
                  }(_pat_let29_0);
                }(_pat_let_tv24);
              }(_pat_let28_0);
            }(_1367_v32);
            _nw206[(new BigNumber(6)).toNumber()] = _module.D2.create_DC7(_dafny.Seq.update(_1349_v17, _module.__default.safeIndex(p2, new BigNumber((_1349_v17).length)), _1331_v0), p1, (_dafny.ZERO).minus(p0), _this.f13);
            _nw206[(new BigNumber(7)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(8)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(9)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(10)).toNumber()] = _module.D2.create_DC7(_dafny.Seq.of(_1331_v0), new BigNumber((_1368_v33).length), p0, _this.f13);
            _nw206[(new BigNumber(11)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(12)).toNumber()] = _module.__default.fm23(p1, _this.f13, globalState);
            _nw206[(new BigNumber(13)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(14)).toNumber()] = _1367_v32;
            _nw206[(new BigNumber(15)).toNumber()] = _module.__default.fm23(p1, _this.f13, globalState);
            _nw206[(new BigNumber(16)).toNumber()] = _module.D2.create_DC7((_1369_v34)[_module.__default.safeIndex(p2, new BigNumber((_1369_v34).length))], p0, p0, _this.f13);
            _nw206[(new BigNumber(17)).toNumber()] = _module.D2.create_DC7(_dafny.Seq.of(_1331_v0, _1331_v0), p2, p1, _this.f13);
            _nw206[(new BigNumber(18)).toNumber()] = _1367_v32;
            _1370_v35 = _nw206;
            let _index251 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1370_v35).length));
            (_1370_v35)[_index251] = _1367_v32;
            let _index252 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1370_v35).length));
            (_1370_v35)[_index252] = _1367_v32;
            (_this).f13 = _this.f13;
          }
        }
      }
      if (_this.f13) {
        let _rhs206 = new BigNumber(-866);
        let _rhs207 = _this.f13;
        let _rhs208 = (p2).minus((p2).minus(p2));
        let _rhs209 = new BigNumber(107);
        r1 = _rhs206;
        r0 = _rhs207;
        r1 = _rhs208;
        r1 = _rhs209;
        r1 = p2;
        let _1373_v36;
        _1373_v36 = _module.D9.create_DC24(_this.f13, _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("nxkq"), _1331_v0), _1331_v0);
        let _source23 = _1373_v36;
        if (_source23.is_DC24) {
          let _1374___mcc_h0 = (_source23).cf36;
          let _1375___mcc_h1 = (_source23).cf37;
          let _1376___mcc_h2 = (_source23).cf38;
          let _1377_cf38 = _1376___mcc_h2;
          let _1378_cf37 = _1375___mcc_h1;
          let _1379_cf36 = _1374___mcc_h0;
          (_this).f13 = _1379_cf36;
          let _1380_v37;
          _1380_v37 = _module.D7.create_DC20(p1, _1331_v0);
          let _1381_v38;
          _1381_v38 = _module.D7.create_DC21(_1380_v37);
          let _1382_v39;
          _1382_v39 = _module.D7.create_DC21(_1380_v37);
          let _1383_v40;
          _1383_v40 = _module.D7.create_DC21(_1382_v39);
          let _1384_v41;
          _1384_v41 = _dafny.Seq.of(_1383_v40, _1383_v40);
          let _1385_v42;
          let _nw207 = new _module.C2();
          _nw207.__ctor(!_dafny.areEqual(_1384_v41, _1384_v41));
          _1385_v42 = _nw207;
          _1385_v42 = _1385_v42;
          let _1386_v43;
          _1386_v43 = _dafny.Seq.of(p1, p1);
          let _1387_v44;
          _1387_v44 = _dafny.MultiSet.fromElements(_1379_cf36);
          r0 = ((_1386_v43)[_module.__default.safeIndex(new BigNumber((_1387_v44).cardinality()), new BigNumber((_1386_v43).length))]).isLessThan(p2);
          let _1388_v45;
          let _nw208 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1388_v45 = _nw208;
          let _1389_v46;
          _1389_v46 = _module.D4.create_DC10(_1347_v16);
          let _index253 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1388_v45).length));
          (_1388_v45)[_index253] = (_1389_v46).dtor_cf20;
          let _index254 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1388_v45).length));
          (_1388_v45)[_index254] = _1347_v16;
        } else if (_source23.is_DC25) {
          let _1390___mcc_h3 = (_source23).cf39;
          let _1391___mcc_h4 = (_source23).cf40;
          let _1392___mcc_h5 = (_source23).cf41;
          let _1393_cf41 = _1392___mcc_h5;
          let _1394_cf40 = _1391___mcc_h4;
          let _1395_cf39 = _1390___mcc_h3;
          let _1396_v47;
          _1396_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7((_dafny.ZERO).minus(p2), p0, _1395_cf39, globalState),p0);
          let _1397_v48;
          let _nw209 = new _module.C2();
          _nw209.__ctor(_module.__default.fm0(_1396_v47, p0, globalState));
          _1397_v48 = _nw209;
          let _1398_v49;
          _1398_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1394_cf40,_1397_v48);
          let _rhs210 = (_1393_cf41) || (_1395_cf39);
          let _rhs211 = (p1).isLessThanOrEqualTo(new BigNumber((_1398_v49).length));
          r0 = _rhs210;
          _1395_cf39 = _rhs211;
          let _1399_v50;
          _1399_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new BigNumber((_dafny.Set.fromElements(_1395_cf39, _1393_cf41)).length));
          let _1400_v51;
          _1400_v51 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1399_v50).length)), new BigNumber(690), (_dafny.ZERO).minus(_1394_cf40));
          let _1401_v52;
          _1401_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1395_cf39,_1395_cf39);
          let _1402_v53;
          _1402_v53 = _dafny.Map.Empty.slice().updateUnsafe((_1400_v51)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1400_v51).length))],_1401_v52);
          let _1403_v54;
          _1403_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1331_v0,_1394_cf40);
          r1 = (new BigNumber(((_1402_v53).update(new BigNumber((_1403_v54).length), _1401_v52)).length)).plus(p2);
          r1 = _module.__default.safeDivisionInt(_1394_cf40, _1394_cf40);
          let _index255 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1347_v16).length));
          (_1347_v16)[_index255] = ((_1395_cf39) ? (p0) : (new BigNumber(664)));
          let _1404_v55;
          _1404_v55 = _module.D0.create_DC1(_1395_cf39, new BigNumber(209));
          let _1405_v56;
          _1405_v56 = _dafny.Set.fromElements(_1331_v0, _1331_v0);
          let _index256 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1347_v16).length));
          let _rhs212 = _module.__default.safeDivisionInt(new BigNumber(((_module.__default.fm47(_1404_v55, p0, globalState)).Union(function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_1405_v56).Elements) {
              let _1406_v57 = _compr_44;
              if ((_1405_v56).contains(_1406_v57)) {
                _coll44.add(_1406_v57);
              }
            }
            return _coll44;
          }())).length), p2);
          let _rhs213 = new BigNumber((_1349_v17).length);
          let _rhs214 = _1399_v50;
          let _lhs132 = _1347_v16;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1347_v16).length));
          r1 = _rhs212;
          _lhs132[_lhs133] = _rhs213;
          _1399_v50 = _rhs214;
        } else {
          let _1407___mcc_h6 = (_source23).cf35;
          let _1408_cf35 = _1407___mcc_h6;
          let _1409_v58;
          let _nw210 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1409_v58 = _nw210;
          let _index257 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          (_1409_v58)[_index257] = ((_this.f13) ? (_this.f13) : (_this.f13));
          let _1410_v59;
          let _nw211 = new _module.C1();
          _nw211.__ctor(_this.f13, _1349_v17, (_this).f14);
          _1410_v59 = _nw211;
          let _1411_v60;
          _1411_v60 = _dafny.Seq.of(_this.f13);
          let _1412_v61;
          _1412_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1410_v59,(_dafny.MultiSet.FromArray(_1411_v60)).Union(_1408_cf35));
          let _1413_v62;
          _1413_v62 = _dafny.MultiSet.fromElements(_1410_v59.f18, _1410_v59.f18);
          let _1414_v63;
          _1414_v63 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f13);
          let _1415_v64;
          _1415_v64 = _dafny.Seq.of(_1349_v17, _1410_v59.f18, _1349_v17);
          let _index258 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          let _rhs215 = (((_1413_v62).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-486)), ((_1416_v0) => function (_1417_i4) {
            return _1416_v0;
          })(_1331_v0)), _module.__default.abs(new BigNumber((_1414_v63).length)))).update((_1415_v64)[_module.__default.safeIndex(p1, new BigNumber((_1415_v64).length))], _module.__default.abs((((_1408_cf35).contains(_this.f13)) ? ((_1408_cf35).get(_this.f13)) : (p0))))).IsDisjointFrom(_1413_v62);
          let _rhs216 = _this.f13;
          let _rhs217 = (((_1410_v59).f17) ? (((_this.f13) ? (_this.f13) : (false))) : (!(_this.f13)));
          let _rhs218 = ((!((_1410_v59).f17) || ((_1410_v59).f17)) ? (_1412_v61) : ((_1412_v61).Merge(_1412_v61)));
          let _lhs134 = _this;
          let _lhs135 = _this;
          let _lhs136 = _1409_v58;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          _lhs134.f13 = _rhs215;
          _lhs135.f13 = _rhs216;
          _lhs136[_lhs137] = _rhs217;
          _1412_v61 = _rhs218;
          let _index259 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length));
          (_1347_v16)[_index259] = p2;
          let _index260 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          let _index261 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length));
          let _index262 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          let _rhs219 = (_1409_v58)[_module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length))];
          let _rhs220 = p0;
          let _rhs221 = ((_dafny.ZERO).minus(p2)).minus(p0);
          let _rhs222 = (_1410_v59).f17;
          let _rhs223 = (_1408_cf35).Difference(_1408_cf35);
          let _lhs138 = _1409_v58;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          let _lhs140 = _1347_v16;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length));
          let _lhs142 = _1409_v58;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1409_v58).length));
          _lhs138[_lhs139] = _rhs219;
          _lhs140[_lhs141] = _rhs220;
          r1 = _rhs221;
          _lhs142[_lhs143] = _rhs222;
          _1408_cf35 = _rhs223;
          let _1418_v65;
          _1418_v65 = _dafny.Set.fromElements(_1349_v17);
          let _1419_v66;
          let _nw212 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _1419_v66 = _nw212;
          let _1420_v67;
          _1420_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1410_v59.f18,(_1410_v59).f17);
          let _index263 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length));
          let _rhs224 = (p0).isLessThanOrEqualTo(_module.__default.safeDivisionInt(p1, new BigNumber((_1411_v60).length)));
          let _rhs225 = function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of (_1420_v67).Keys.Elements) {
              let _1421_v68 = _compr_45;
              if ((_1420_v67).contains(_1421_v68)) {
                _coll45.add(_1421_v68);
              }
            }
            return _coll45;
          }();
          let _rhs226 = _1331_v0;
          let _rhs227 = (_this).f14;
          let _rhs228 = (_1410_v59).fm7(new BigNumber(517), ((_1347_v16)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length))]).multipliedBy(p2), (_1410_v59).f17, globalState);
          let _lhs144 = _1347_v16;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1347_v16).length));
          r0 = _rhs224;
          _1418_v65 = _rhs225;
          _1331_v0 = _rhs226;
          _1419_v66 = _rhs227;
          _lhs144[_lhs145] = _rhs228;
          r0 = !(true);
        }
        let _1422_v69;
        let _nw213 = new _module.C0();
        _nw213.__ctor(p0);
        _1422_v69 = _nw213;
        _1349_v17 = _1349_v17;
      } else {
        let _1423_v70;
        _1423_v70 = _dafny.Set.fromElements(_this.f13);
        let _1424_v71;
        _1424_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,new BigNumber((_1423_v70).length));
        let _1425_v72;
        _1425_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v71,!(_this.f13));
        let _1426_v73;
        _1426_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-410),new BigNumber((_1425_v72).length));
        let _1427_v74;
        _1427_v74 = _dafny.Seq.of(new BigNumber(-185));
        let _1428_v75;
        _1428_v75 = _dafny.Set.fromElements(p0, new BigNumber((_1427_v74).length));
        let _1429_v76;
        _1429_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_1426_v73).contains(p0)) ? ((_1426_v73).get(p0)) : (new BigNumber((_1428_v75).length))));
        let _1430_v77;
        _1430_v77 = _dafny.Set.fromElements(_dafny.Seq.of(_this.f13));
        _1429_v76 = (_1429_v76).update(p0, _module.__default.safeModuloInt(new BigNumber((_1430_v77).length), p0));
        r1 = _module.__default.safeModuloInt(p2, p2);
        let _1431_v78;
        _1431_v78 = _module.D9.create_DC25(_this.f13, (_dafny.ZERO).minus(p0), _this.f13);
        let _1432_v79;
        _1432_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f13,_this.f13);
        (_this).f13 = !(!((_module.__default.fm20(_this.f13, p0, (_1431_v78).dtor_cf41, globalState)).equals(_1432_v79)) || (false));
        (_this).f13 = !(_this.f13);
        let _1433_v80;
        let _init34 = ((_1434_v0) => function (_1435_i5) {
          return _1434_v0;
        })(_1331_v0);
        let _nw214 = Array((new BigNumber(11)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw214.length); _i0_34++) {
          _nw214[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1433_v80 = _nw214;
        let _index264 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1433_v80).length));
        (_1433_v80)[_index264] = _1331_v0;
        let _index265 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1433_v80).length));
        let _rhs229 = new BigNumber(92);
        let _rhs230 = _1331_v0;
        let _lhs146 = _1433_v80;
        let _lhs147 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1433_v80).length));
        r1 = _rhs229;
        _lhs146[_lhs147] = _rhs230;
      }
      let _1436_v81;
      let _nw215 = Array((new BigNumber(9)).toNumber()).fill(false);
      _1436_v81 = _nw215;
      let _1437_v82;
      _1437_v82 = _dafny.MultiSet.fromElements(_1436_v81);
      r0 = (_1437_v82).IsSubsetOf(_1437_v82);
      let _1438_v83;
      _1438_v83 = _module.D1.create_DC5(_1349_v17, _1436_v81, p0);
      r1 = (_1438_v83).dtor_cf10;
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
