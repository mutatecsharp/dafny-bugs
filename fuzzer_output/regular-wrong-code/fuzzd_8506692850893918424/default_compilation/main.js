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
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))));
    };
    static fm2(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('s'.codePointAt(0));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(!(false))).Intersect(_dafny.Set.fromElements(!(false)))).Union(_dafny.Set.fromElements(true));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return (new BigNumber(-314)).isLessThanOrEqualTo((new BigNumber(548)).plus(new BigNumber((_dafny.Seq.of(new BigNumber(412))).length)));
    };
    static fm5(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(421), new BigNumber(-17))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, !(false), false)).length), new BigNumber((_dafny.Seq.of(false)).length)))));
    };
    static fm6(p0, p1, globalState) {
      return (new BigNumber(100)).plus((_module.D1.create_DC3(new BigNumber(339), !(!(!(!(true)))))).dtor_cf6);
    };
    static fm7(p0, p1, globalState) {
      let _source0 = _module.D3.create_DC8(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-463), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("u")).length)), new BigNumber(851), new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("cdurb"), _dafny.Seq.UnicodeFromString("amh"))).length), new BigNumber(331))).cardinality()));
      if (_source0.is_DC8) {
        let _0___mcc_h0 = (_source0).cf13;
        let _1_cf13 = _0___mcc_h0;
        return _module.D1.create_DC3(_1_cf13, false);
      } else if (_source0.is_DC7) {
        let _2___mcc_h1 = (_source0).cf12;
        let _3_cf12 = _2___mcc_h1;
        return _module.D1.create_DC3(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()), false);
      } else {
        let _4___mcc_h2 = (_source0).cf14;
        let _5_cf14 = _4___mcc_h2;
        return _module.D1.create_DC3(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(false))).length))).cardinality()), false);
      }
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber(767), new BigNumber((_dafny.Seq.UnicodeFromString("nux")).length)), (new BigNumber((_dafny.Set.fromElements(new BigNumber(-871), new BigNumber((_dafny.Seq.UnicodeFromString("rqqxre")).length))).length)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length)));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D2.create_DC6();
      if (_source1.is_DC5) {
        let _6___mcc_h0 = (_source1).cf9;
        let _7___mcc_h1 = (_source1).cf10;
        let _8___mcc_h2 = (_source1).cf11;
        let _9_cf11 = _8___mcc_h2;
        let _10_cf10 = _7___mcc_h1;
        let _11_cf9 = _6___mcc_h0;
        return _dafny.Seq.of(_module.D0.create_DC0(new _dafny.CodePoint('y'.codePointAt(0))), _module.D0.create_DC0(new _dafny.CodePoint('j'.codePointAt(0))), _module.D0.create_DC0(new _dafny.CodePoint('w'.codePointAt(0))), _module.D0.create_DC0(new _dafny.CodePoint('h'.codePointAt(0))));
      } else if (_source1.is_DC6) {
        return _dafny.Seq.of(_module.D0.create_DC0(new _dafny.CodePoint('b'.codePointAt(0))), _module.D0.create_DC0(new _dafny.CodePoint('p'.codePointAt(0))));
      } else {
        let _12___mcc_h3 = (_source1).cf8;
        let _13_cf8 = _12___mcc_h3;
        return _dafny.Seq.of(_module.D0.create_DC0(new _dafny.CodePoint('y'.codePointAt(0))), _module.D0.create_DC0(new _dafny.CodePoint('r'.codePointAt(0))));
      }
    };
    static m0(globalState) {
      let r0 = false;
      let r1 = _dafny.Set.Empty;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _14_v0;
      _14_v0 = new BigNumber(506);
      let _15_v1;
      _15_v1 = _dafny.Seq.UnicodeFromString("skwjs");
      let _16_v2;
      _16_v2 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_14_v0, new BigNumber((_15_v1).length)),_dafny.Seq.UnicodeFromString("k"));
      let _17_v3;
      _17_v3 = _dafny.Seq.of(_15_v1);
      _16_v2 = (_16_v2).update((_14_v0).multipliedBy(new BigNumber(61)), _dafny.Seq.Concat((_17_v3)[_module.__default.safeIndex(_14_v0, new BigNumber((_17_v3).length))], _15_v1));
      let _18_v4;
      let _nw0 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
      _18_v4 = _nw0;
      let _19_v5;
      _19_v5 = true;
      let _20_v6;
      _20_v6 = _module.D1.create_DC3((_dafny.ZERO).minus(_14_v0), _19_v5);
      let _21_v7;
      _21_v7 = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_dafny.Set.fromElements(_20_v6));
      let _pat_let_tv0 = _14_v0;
      let _index0 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_18_v4).length));
      (_18_v4)[_index0] = (((_21_v7).contains(_19_v5)) ? ((_21_v7).get(_19_v5)) : (_dafny.Set.fromElements(_20_v6, function (_pat_let0_0) {
        return function (_22_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_23_dt__update_hcf6_h0) {
              return _module.D1.create_DC3(_23_dt__update_hcf6_h0, (_22_dt__update__tmp_h0).dtor_cf7);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_20_v6), _20_v6)));
      let _24_v8;
      _24_v8 = _dafny.Set.fromElements(_20_v6);
      let _index1 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_18_v4).length));
      (_18_v4)[_index1] = _24_v8;
      let _25_i0;
      _25_i0 = _dafny.ZERO;
      L0: {
        while ((_19_v5) === (_19_v5)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_25_i0)) {
              break L0;
            }
            _25_i0 = (_25_i0).plus(_dafny.ONE);
            _14_v0 = (_dafny.ZERO).minus(_14_v0);
            let _26_v9;
            let _init0 = ((_27_v0) => function (_28_i1) {
              return _module.__default.safeDivisionInt(_28_i1, _27_v0);
            })(_14_v0);
            let _nw1 = Array((new BigNumber(29)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
              _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _26_v9 = _nw1;
            let _index2 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_26_v9).length));
            (_26_v9)[_index2] = _14_v0;
            let _index3 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_26_v9).length));
            (_26_v9)[_index3] = _14_v0;
            let _29_v10;
            _29_v10 = new _dafny.CodePoint('k'.codePointAt(0));
            let _30_v11;
            _30_v11 = _module.D0.create_DC0(_29_v10);
            let _31_v12;
            _31_v12 = _dafny.Seq.of(_30_v11);
            let _32_v13;
            _32_v13 = _dafny.Seq.of((_26_v9)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_26_v9).length))]);
            let _33_v14;
            _33_v14 = _dafny.Seq.of(true, _19_v5, _19_v5, !(_19_v5));
            let _34_v15;
            _34_v15 = _dafny.Seq.of((_26_v9)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_26_v9).length))], (_32_v13)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(650),_19_v5)).length), new BigNumber((_32_v13).length))], new BigNumber((_33_v14).length), _14_v0, new BigNumber((_15_v1).length));
            let _rhs0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_31_v12, _module.__default.fm9(_14_v0, (_26_v9)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_26_v9).length))], _19_v5, _19_v5, globalState)), _31_v12);
            let _rhs1 = _module.__default.fm4(_15_v1, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), ((_35_v1) => function (_36_i2) {
              return new BigNumber((_35_v1).length);
            })(_15_v1)), _32_v13), _19_v5, _19_v5, globalState);
            _31_v12 = _rhs0;
            _19_v5 = _rhs1;
            r0 = _19_v5;
          }
        }
      }
      let _37_v16;
      _37_v16 = _dafny.Seq.of(_14_v0);
      (globalState).f6 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(706), _14_v0), (_14_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(476)), function (_38_i3) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _37_v16, _19_v5, _19_v5, globalState)),_14_v0)).length))));
      let _39_v17;
      let _nw2 = new _module.C0();
      _nw2.__ctor();
      _39_v17 = _nw2;
      let _40_v18;
      _40_v18 = _dafny.Map.Empty.slice().updateUnsafe(_14_v0,_39_v17);
      let _41_v19;
      _41_v19 = _dafny.MultiSet.fromElements(_40_v18, _40_v18);
      let _42_v20;
      _42_v20 = _dafny.Set.fromElements(_19_v5);
      let _43_v21;
      _43_v21 = _dafny.Map.Empty.slice().updateUnsafe((((_41_v19).contains(_40_v18)) ? ((_41_v19).get(_40_v18)) : (_module.__default.fm6(_14_v0, _19_v5, globalState))),(_42_v20).IsProperSubsetOf(_dafny.Set.fromElements(!(_19_v5))));
      if ((((_43_v21).contains((_14_v0).multipliedBy(_14_v0))) ? ((_43_v21).get((_14_v0).multipliedBy(_14_v0))) : (true))) {
        let _44_v22;
        let _nw3 = new _module.C0();
        _nw3.__ctor();
        _44_v22 = _nw3;
        let _45_v23;
        _45_v23 = _module.D0.create_DC1((_14_v0).isEqualTo(_14_v0), _14_v0, _19_v5, (_14_v0).multipliedBy(_14_v0));
        _45_v23 = _45_v23;
        let _46_v24;
        _46_v24 = _dafny.Seq.of(false, _19_v5);
        let _47_v25;
        _47_v25 = _dafny.Seq.of(_dafny.Seq.of(_19_v5, _19_v5), _dafny.Seq.of(false));
        let _48_v26;
        _48_v26 = _dafny.Seq.of(_39_v17);
        let _rhs2 = !_dafny.Seq.contains(_47_v25, _46_v24);
        let _rhs3 = _19_v5;
        let _rhs4 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(883)), function (_49_i4) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })).length), new BigNumber((_48_v26).length));
        let _lhs0 = globalState;
        _19_v5 = _rhs2;
        r0 = _rhs3;
        _lhs0.f6 = _rhs4;
        let _50_v27;
        _50_v27 = _dafny.MultiSet.fromElements(_14_v0, _14_v0, _14_v0);
        let _51_v28;
        _51_v28 = _dafny.MultiSet.fromElements(_19_v5);
        let _52_v29;
        _52_v29 = new _dafny.CodePoint('g'.codePointAt(0));
        let _53_v30;
        _53_v30 = _dafny.Map.Empty.slice().updateUnsafe(_52_v29,!(true));
        let _54_v31;
        let _nw4 = Array((new BigNumber(23)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _19_v5;
        _nw4[(_dafny.ONE).toNumber()] = (_20_v6).dtor_cf7;
        _nw4[(new BigNumber(2)).toNumber()] = _dafny.areEqual(_15_v1, _15_v1);
        _nw4[(new BigNumber(3)).toNumber()] = ((!(_19_v5)) ? (true) : (_19_v5));
        _nw4[(new BigNumber(4)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(5)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(6)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(7)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(8)).toNumber()] = !(_14_v0).isEqualTo(_14_v0);
        _nw4[(new BigNumber(9)).toNumber()] = (new BigNumber((_37_v16).length)).isLessThan(_14_v0);
        _nw4[(new BigNumber(10)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(11)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(12)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(13)).toNumber()] = !((_19_v5) || (false));
        _nw4[(new BigNumber(14)).toNumber()] = (new BigNumber((_42_v20).length)).isLessThan((_dafny.ZERO).minus(_14_v0));
        _nw4[(new BigNumber(15)).toNumber()] = (_50_v27).IsDisjointFrom(_50_v27);
        _nw4[(new BigNumber(16)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(17)).toNumber()] = (_51_v28).IsProperSubsetOf(_51_v28);
        _nw4[(new BigNumber(18)).toNumber()] = _19_v5;
        _nw4[(new BigNumber(19)).toNumber()] = true;
        _nw4[(new BigNumber(20)).toNumber()] = (((_53_v30).contains(_52_v29)) ? ((_53_v30).get(_52_v29)) : (_19_v5));
        _nw4[(new BigNumber(21)).toNumber()] = !(_19_v5);
        _nw4[(new BigNumber(22)).toNumber()] = _19_v5;
        _54_v31 = _nw4;
        let _index4 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        (_54_v31)[_index4] = _19_v5;
        let _index5 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        (_54_v31)[_index5] = !(_19_v5) || (_19_v5);
        let _55_v32;
        _55_v32 = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_52_v29);
        let _index6 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        let _index7 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        let _rhs5 = (_54_v31)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length))];
        let _rhs6 = (((_dafny.ZERO).minus(_14_v0)).minus(_14_v0)).isLessThan(_14_v0);
        let _rhs7 = (_14_v0).isLessThanOrEqualTo(new BigNumber((_15_v1).length));
        let _rhs8 = _dafny.Seq.update(_15_v1, _module.__default.safeIndex(_14_v0, new BigNumber((_15_v1).length)), _52_v29);
        let _rhs9 = _55_v32;
        let _lhs1 = _54_v31;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        let _lhs3 = _54_v31;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_54_v31).length));
        r0 = _rhs5;
        _lhs1[_lhs2] = _rhs6;
        _lhs3[_lhs4] = _rhs7;
        r2 = _rhs8;
        _55_v32 = _rhs9;
      } else {
        (globalState).f7 = (_20_v6).dtor_cf6;
        (globalState).f6 = (_dafny.ZERO).minus((_37_v16)[_module.__default.safeIndex(_14_v0, new BigNumber((_37_v16).length))]);
        let _56_v33;
        let _init1 = ((_57_v0) => function (_58_i5) {
          return (_58_i5).minus(_57_v0);
        })(_14_v0);
        let _nw5 = Array((new BigNumber(21)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
          _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _56_v33 = _nw5;
        let _59_v34;
        _59_v34 = _dafny.MultiSet.fromElements(_56_v33, _56_v33);
        if ((_dafny.MultiSet.fromElements(_56_v33)).equals(_59_v34)) {
          let _60_v35;
          _60_v35 = _dafny.Seq.of(_19_v5, _19_v5);
          _60_v35 = _dafny.Seq.Concat(_60_v35, _60_v35);
          _19_v5 = ((_14_v0).minus(_14_v0)).isLessThan(_14_v0);
          (globalState).f8 = (new BigNumber(245)).isLessThanOrEqualTo((new BigNumber((_60_v35).length)).multipliedBy(_14_v0));
          let _index8 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_56_v33).length));
          (_56_v33)[_index8] = (_14_v0).minus(_14_v0);
          let _61_v36;
          _61_v36 = _dafny.Map.Empty.slice().updateUnsafe(_56_v33,_module.__default.safeDivisionInt(_14_v0, _14_v0));
          let _62_v37;
          _62_v37 = _dafny.Seq.of(_56_v33);
          let _63_v38;
          _63_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_62_v37).length),_14_v0);
          let _64_v39;
          _64_v39 = _dafny.Set.fromElements(_63_v38, _63_v38);
          let _65_v40;
          _65_v40 = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_19_v5);
          let _66_v41;
          _66_v41 = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_39_v17);
          let _index9 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_56_v33).length));
          let _rhs10 = _19_v5;
          let _rhs11 = (((_61_v36).contains(_56_v33)) ? ((_61_v36).get(_56_v33)) : (_14_v0));
          let _rhs12 = _14_v0;
          let _rhs13 = _19_v5;
          let _rhs14 = (_module.D0.create_DC1(_19_v5, new BigNumber((_64_v39).length), (((_65_v40).contains(_module.__default.fm4(_15_v1, _dafny.Seq.of(_14_v0), _19_v5, _19_v5, globalState))) ? ((_65_v40).get(_module.__default.fm4(_15_v1, _dafny.Seq.of(_14_v0), _19_v5, _19_v5, globalState))) : (_19_v5)), new BigNumber((_66_v41).length))).dtor_cf1;
          let _lhs5 = _56_v33;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_56_v33).length));
          let _lhs7 = globalState;
          let _lhs8 = globalState;
          r0 = _rhs10;
          _lhs5[_lhs6] = _rhs11;
          _lhs7.f6 = _rhs12;
          r0 = _rhs13;
          _lhs8.f8 = _rhs14;
          _60_v35 = _dafny.Seq.Concat(_60_v35, _60_v35);
        } else {
          (globalState).f6 = _14_v0;
          let _67_v42;
          _67_v42 = _dafny.Seq.of(_39_v17, _39_v17);
          _39_v17 = (_67_v42)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(473)), function (_68_i6) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length), new BigNumber((_67_v42).length))];
          let _69_v43;
          _69_v43 = _dafny.Seq.of(_19_v5, _19_v5);
          let _70_v44;
          _70_v44 = _dafny.MultiSet.fromElements(_14_v0);
          let _71_v45;
          _71_v45 = _dafny.Map.Empty.slice().updateUnsafe(_69_v43,(_70_v44).IsDisjointFrom(_70_v44));
          let _rhs15 = _module.__default.safeDivisionInt(_14_v0, _module.__default.safeDivisionInt(new BigNumber(-380), _14_v0));
          let _rhs16 = (new BigNumber((_dafny.Seq.Concat(_15_v1, _15_v1)).length)).multipliedBy(_14_v0);
          let _rhs17 = _71_v45;
          let _lhs9 = globalState;
          let _lhs10 = globalState;
          _lhs9.f3 = _rhs15;
          _lhs10.f3 = _rhs16;
          _71_v45 = _rhs17;
          let _72_v46;
          let _nw6 = new _module.C0();
          _nw6.__ctor();
          _72_v46 = _nw6;
          (globalState).f3 = _14_v0;
        }
        let _73_v47;
        let _nw7 = Array((new BigNumber(28)).toNumber()).fill([]);
        _73_v47 = _nw7;
        let _74_v48;
        let _init2 = ((_75_v5) => function (_76_i7) {
          return _75_v5;
        })(_19_v5);
        let _nw8 = Array((new BigNumber(17)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw8.length); _i0_2++) {
          _nw8[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _74_v48 = _nw8;
        let _index10 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_73_v47).length));
        (_73_v47)[_index10] = _74_v48;
        let _index11 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_73_v47).length));
        (_73_v47)[_index11] = _74_v48;
        let _77_v49;
        _77_v49 = _module.D0.create_DC1(_19_v5, new BigNumber(34), !(_19_v5), (_dafny.ZERO).minus(new BigNumber(((_42_v20).Union(_42_v20)).length)));
        let _source2 = _77_v49;
        if (_source2.is_DC1) {
          let _78___mcc_h0 = (_source2).cf1;
          let _79___mcc_h1 = (_source2).cf2;
          let _80___mcc_h2 = (_source2).cf3;
          let _81___mcc_h3 = (_source2).cf4;
          let _82_cf4 = _81___mcc_h3;
          let _83_cf3 = _80___mcc_h2;
          let _84_cf2 = _79___mcc_h1;
          let _85_cf1 = _78___mcc_h0;
          r2 = _15_v1;
          _74_v48 = (_73_v47)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_73_v47).length))];
          (globalState).f3 = new BigNumber((_dafny.Seq.UnicodeFromString("vspevqskp")).length);
          let _86_v50;
          _86_v50 = _dafny.Map.Empty.slice().updateUnsafe(_83_cf3,_82_cf4);
          _16_v2 = (_16_v2).update((new BigNumber((_86_v50).length)).multipliedBy(_14_v0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), function (_87_i8) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }));
        } else {
          let _88___mcc_h4 = (_source2).cf0;
          let _89_cf0 = _88___mcc_h4;
          (globalState).f3 = (_37_v16)[_module.__default.safeIndex(new BigNumber((_42_v20).length), new BigNumber((_37_v16).length))];
          let _90_v51;
          _90_v51 = _dafny.MultiSet.fromElements(_19_v5);
          let _91_v52;
          _91_v52 = _dafny.Set.fromElements(_90_v51, _90_v51, _90_v51);
          let _92_v53;
          _92_v53 = _dafny.Map.Empty.slice().updateUnsafe(true,_91_v52);
          _91_v52 = (((_92_v53).contains(_module.__default.fm4(_dafny.Seq.UnicodeFromString("x"), _37_v16, !(!(_19_v5)), !(_19_v5), globalState))) ? ((_92_v53).get(_module.__default.fm4(_dafny.Seq.UnicodeFromString("x"), _37_v16, !(!(_19_v5)), !(_19_v5), globalState))) : (_91_v52));
          (globalState).f8 = _19_v5;
          let _93_v54;
          let _nw9 = Array((new BigNumber(4)).toNumber()).fill([]);
          _93_v54 = _nw9;
          let _94_v55;
          let _init3 = function (_95_i9) {
            return _dafny.Seq.of(false);
          };
          let _nw10 = Array((new BigNumber(19)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw10.length); _i0_3++) {
            _nw10[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _94_v55 = _nw10;
          let _index12 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_93_v54).length));
          (_93_v54)[_index12] = _94_v55;
          let _96_v56;
          _96_v56 = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_15_v1, _37_v16, _19_v5, _19_v5, globalState),_14_v0));
          let _97_v57;
          _97_v57 = _dafny.Seq.of(_94_v55, _94_v55);
          let _98_v58;
          _98_v58 = _dafny.Seq.of(_94_v55, (_97_v57)[_module.__default.safeIndex((_dafny.ZERO).minus(_14_v0), new BigNumber((_97_v57).length))], _94_v55);
          let _index13 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_93_v54).length));
          let _rhs18 = (_97_v57)[_module.__default.safeIndex(_14_v0, new BigNumber((_97_v57).length))];
          let _rhs19 = (_module.__default.fm6(_14_v0, _19_v5, globalState)).multipliedBy(_module.__default.fm6(_14_v0, _19_v5, globalState));
          let _rhs20 = _96_v56;
          let _lhs11 = _93_v54;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_93_v54).length));
          let _lhs13 = globalState;
          _lhs11[_lhs12] = _rhs18;
          _lhs13.f6 = _rhs19;
          _96_v56 = _rhs20;
        }
      }
      let _99_v59;
      _99_v59 = _dafny.Seq.of(_42_v20, _42_v20);
      let _100_i10;
      _100_i10 = _dafny.ZERO;
      L1: {
        while ((new BigNumber((_37_v16).length)).isLessThanOrEqualTo(new BigNumber(((_99_v59)[_module.__default.safeIndex(_14_v0, new BigNumber((_99_v59).length))]).length))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_100_i10)) {
              break L1;
            }
            _100_i10 = (_100_i10).plus(_dafny.ONE);
            let _101_v60;
            let _init4 = ((_102_v0) => function (_103_i11) {
              return (_103_i11).multipliedBy(_102_v0);
            })(_14_v0);
            let _nw11 = Array((new BigNumber(7)).toNumber());
            for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw11.length); _i0_4++) {
              _nw11[_i0_4] = _init4(new BigNumber(_i0_4));
            }
            _101_v60 = _nw11;
            let _104_v61;
            _104_v61 = _dafny.MultiSet.fromElements(_101_v60, _101_v60, _101_v60, _101_v60);
            let _rhs21 = (_module.D2.create_DC4(_101_v60)).dtor_cf8;
            let _rhs22 = _14_v0;
            let _rhs23 = _19_v5;
            let _rhs24 = _19_v5;
            let _rhs25 = new BigNumber(((_104_v61).update(_101_v60, _module.__default.abs((_14_v0).minus(_14_v0)))).cardinality());
            let _lhs14 = globalState;
            let _lhs15 = globalState;
            let _lhs16 = globalState;
            _101_v60 = _rhs21;
            _14_v0 = _rhs22;
            _lhs14.f8 = _rhs23;
            _lhs15.f8 = _rhs24;
            _lhs16.f6 = _rhs25;
            if (_19_v5) {
              let _105_v62;
              _105_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(292),_14_v0);
              let _106_v63;
              _106_v63 = new _dafny.CodePoint('p'.codePointAt(0));
              let _107_v64;
              _107_v64 = _dafny.MultiSet.fromElements(_106_v63);
              let _108_v65;
              let _nw12 = Array((new BigNumber(21)).toNumber());
              _nw12[(_dafny.ZERO).toNumber()] = _19_v5;
              _nw12[(_dafny.ONE).toNumber()] = _19_v5;
              _nw12[(new BigNumber(2)).toNumber()] = _module.__default.fm4(_dafny.Seq.UnicodeFromString("owms"), _37_v16, _19_v5, _19_v5, globalState);
              _nw12[(new BigNumber(3)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(4)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(5)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(6)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(7)).toNumber()] = (_14_v0).isLessThan(new BigNumber((_105_v62).length));
              _nw12[(new BigNumber(8)).toNumber()] = false;
              _nw12[(new BigNumber(9)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(10)).toNumber()] = (_107_v64).IsProperSubsetOf(_107_v64);
              _nw12[(new BigNumber(11)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(12)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(13)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(14)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(15)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(16)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(17)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(18)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(19)).toNumber()] = _19_v5;
              _nw12[(new BigNumber(20)).toNumber()] = _19_v5;
              _108_v65 = _nw12;
              let _index14 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_108_v65).length));
              (_108_v65)[_index14] = ((_19_v5) ? (_19_v5) : (true));
              let _index15 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_108_v65).length));
              (_108_v65)[_index15] = _19_v5;
              _39_v17 = _39_v17;
              let _109_v66;
              let _nw13 = Array((new BigNumber(23)).toNumber());
              _109_v66 = _nw13;
              let _index16 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_109_v66).length));
              (_109_v66)[_index16] = _39_v17;
              let _110_v67;
              _110_v67 = _dafny.Seq.of(_19_v5, (_108_v65)[_module.__default.safeIndex(new BigNumber(836), new BigNumber((_108_v65).length))]);
              let _111_v68;
              _111_v68 = _dafny.MultiSet.fromElements(_19_v5, false);
              let _112_v69;
              _112_v69 = _dafny.Set.fromElements(new BigNumber((_110_v67).length), new BigNumber((_111_v68).cardinality()), _14_v0);
              let _index17 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_109_v66).length));
              let _rhs26 = _39_v17;
              let _rhs27 = _39_v17;
              let _rhs28 = ((new BigNumber((_112_v69).length)).multipliedBy(_14_v0)).plus(_14_v0);
              let _rhs29 = (new BigNumber(708)).multipliedBy(new BigNumber((_43_v21).length));
              let _lhs17 = _109_v66;
              let _lhs18 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_109_v66).length));
              let _lhs19 = globalState;
              let _lhs20 = globalState;
              _39_v17 = _rhs26;
              _lhs17[_lhs18] = _rhs27;
              _lhs19.f3 = _rhs28;
              _lhs20.f3 = _rhs29;
              (globalState).f6 = _14_v0;
              let _113_v70;
              let _nw14 = new _module.C0();
              _nw14.__ctor();
              _113_v70 = _nw14;
            } else {
              (globalState).f8 = _19_v5;
              (globalState).f3 = _14_v0;
              let _114_v71;
              _114_v71 = _dafny.Set.fromElements(_14_v0, (_dafny.ZERO).minus(_14_v0));
              _19_v5 = _dafny.Seq.contains(_dafny.Seq.Concat(_37_v16, _37_v16), new BigNumber((_114_v71).length));
              let _115_v72;
              _115_v72 = _module.D0.create_DC1(false, (_dafny.ZERO).minus(_14_v0), !(_19_v5), _14_v0);
              let _116_v73;
              _116_v73 = _module.D2.create_DC5(_14_v0, _14_v0, _115_v72);
              let _index18 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_101_v60).length));
              (_101_v60)[_index18] = _module.__default.safeModuloInt(_14_v0, (_116_v73).dtor_cf9);
              let _index19 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_101_v60).length));
              (_101_v60)[_index19] = new BigNumber((_dafny.Seq.UnicodeFromString("poflqgduh")).length);
              let _117_v74;
              let _nw15 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
              _117_v74 = _nw15;
              let _118_v75;
              _118_v75 = _dafny.Map.Empty.slice().updateUnsafe(false,_117_v74);
              let _119_v76;
              _119_v76 = _dafny.Seq.of(_118_v75, _118_v75);
              let _120_v77;
              _120_v77 = _module.D3.create_DC7(_118_v75);
              let _121_v78;
              _121_v78 = _dafny.Seq.of(_19_v5);
              let _122_v79;
              let _nw16 = Array((new BigNumber(20)).toNumber());
              _nw16[(_dafny.ZERO).toNumber()] = _118_v75;
              _nw16[(_dafny.ONE).toNumber()] = _118_v75;
              _nw16[(new BigNumber(2)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(3)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(4)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(5)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-688)), function (_123_i12) {
                return new _dafny.CodePoint('k'.codePointAt(0));
              }), _dafny.Seq.of(new BigNumber((_43_v21).length), _14_v0), _19_v5, _19_v5, globalState),_117_v74);
              _nw16[(new BigNumber(7)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_117_v74);
              _nw16[(new BigNumber(9)).toNumber()] = (_118_v75).Merge((_119_v76)[_module.__default.safeIndex(_14_v0, new BigNumber((_119_v76).length))]);
              _nw16[(new BigNumber(10)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(11)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(12)).toNumber()] = (_120_v77).dtor_cf12;
              _nw16[(new BigNumber(13)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(14)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(15)).toNumber()] = _118_v75;
              _nw16[(new BigNumber(16)).toNumber()] = (_118_v75).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_117_v74));
              _nw16[(new BigNumber(17)).toNumber()] = (_118_v75).Merge(_118_v75);
              _nw16[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_121_v78)[_module.__default.safeIndex((_101_v60)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_101_v60).length))], new BigNumber((_121_v78).length))],_117_v74);
              _nw16[(new BigNumber(19)).toNumber()] = _118_v75;
              _122_v79 = _nw16;
              let _index20 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_122_v79).length));
              (_122_v79)[_index20] = (_118_v75).Merge(_dafny.Map.Empty.slice().updateUnsafe(_19_v5,_117_v74));
              let _index21 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_122_v79).length));
              (_122_v79)[_index21] = _118_v75;
            }
            let _124_v80;
            _124_v80 = _module.D0.create_DC1(true, _14_v0, _19_v5, new BigNumber(848));
            (globalState).f7 = _module.__default.safeModuloInt((_124_v80).dtor_cf2, new BigNumber(34));
            let _125_v81;
            let _nw17 = new _module.C0();
            _nw17.__ctor();
            _125_v81 = _nw17;
          }
        }
      }
      r0 = _19_v5;
      r1 = function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_14_v0,_14_v0)).length)), _37_v16)).Elements) {
          let _126_v82 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_14_v0,_14_v0)).length)), _37_v16), _126_v82)) {
            _coll0.add(_module.__default.safeDivisionInt(_126_v82, new BigNumber(669)));
          }
        }
        return _coll0;
      }();
      let _127_v83;
      _127_v83 = _dafny.Map.Empty.slice().updateUnsafe(_19_v5,_14_v0);
      r2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm1(new BigNumber((_127_v83).length), _19_v5, globalState), _15_v1), _15_v1);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _128_v0;
      _128_v0 = new BigNumber(-183);
      let _129_v1;
      _129_v1 = _dafny.MultiSet.fromElements(_128_v0);
      let _130_v2;
      let _nw18 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _130_v2 = _nw18;
      let _131_v3;
      _131_v3 = _dafny.Seq.UnicodeFromString("cqq");
      let _132_v4;
      _132_v4 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rwijw"), _131_v3);
      let _133_globalState;
      let _nw19 = new _module.GlobalState();
      _nw19.__ctor(_129_v1, _130_v2, false, new BigNumber(434), new BigNumber(588), new BigNumber(870), new BigNumber(77), new BigNumber(-390), false, true, _132_v4, true, new BigNumber(980), new BigNumber(718));
      _133_globalState = _nw19;
      let _134_v5;
      _134_v5 = false;
      let _135_v6;
      _135_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_128_v0,!(_134_v5)),_128_v0);
      let _136_v7;
      _136_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-421),false);
      _135_v6 = (_135_v6).update(_136_v7, (_dafny.ZERO).minus(_128_v0));
      let _137_v8;
      let _nw20 = new _module.C0();
      _nw20.__ctor();
      _137_v8 = _nw20;
      let _138_v9;
      let _nw21 = Array((new BigNumber(28)).toNumber()).fill(false);
      _138_v9 = _nw21;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_138_v9).length))) {
        let _139_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_139_i0)) && ((_139_i0).isLessThan(new BigNumber((_138_v9).length))))) {
          (_138_v9)[(_139_i0)] = _134_v5;
        }
      }
      let _140_v10;
      let _141_v11;
      let _142_v12;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = _module.__default.m0(_133_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _140_v10 = _out0;
      _141_v11 = _out1;
      _142_v12 = _out2;
      let _143_i1;
      _143_i1 = _dafny.ZERO;
      L2: {
        while (_140_v10) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_143_i1)) {
              break L2;
            }
            _143_i1 = (_143_i1).plus(_dafny.ONE);
            _140_v10 = !(!(_134_v5)) || (!((_128_v0).isLessThan(_128_v0)));
            let _144_v13;
            let _145_v14;
            let _146_v15;
            let _out3;
            let _out4;
            let _out5;
            let _outcollector1 = _module.__default.m0(_133_globalState);
            _out3 = _outcollector1[0];
            _out4 = _outcollector1[1];
            _out5 = _outcollector1[2];
            _144_v13 = _out3;
            _145_v14 = _out4;
            _146_v15 = _out5;
            (_133_globalState).f3 = new BigNumber((_dafny.Seq.UnicodeFromString("tiuwpcrys")).length);
            let _147_v16;
            _147_v16 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_142_v12);
            _142_v12 = _dafny.Seq.Concat(_146_v15, (((_147_v16).contains(_128_v0)) ? ((_147_v16).get(_128_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), function (_148_i2) {
              return new _dafny.CodePoint('k'.codePointAt(0));
            }))));
          }
        }
      }
      let _149_i3;
      _149_i3 = _dafny.ZERO;
      L3: {
        while ((((_136_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(_128_v0,true))) ? (_140_v10) : (_140_v10))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_149_i3)) {
              break L3;
            }
            _149_i3 = (_149_i3).plus(_dafny.ONE);
            (_133_globalState).f3 = _128_v0;
            let _150_v17;
            let _nw22 = new _module.C0();
            _nw22.__ctor();
            _150_v17 = _nw22;
            _131_v3 = _dafny.Seq.Concat(_131_v3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-155)), function (_151_i4) {
              return (_module.D0.create_DC0(new _dafny.CodePoint('x'.codePointAt(0)))).dtor_cf0;
            }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-596)), function (_152_i5) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })));
            let _153_v18;
            _153_v18 = _dafny.Map.Empty.slice().updateUnsafe(_134_v5,(_134_v5) === (_140_v10));
            (_133_globalState).f6 = new BigNumber((_153_v18).length);
          }
        }
      }
      let _154_i6;
      _154_i6 = _dafny.ZERO;
      L4: {
        while (_134_v5) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_154_i6)) {
              break L4;
            }
            _154_i6 = (_154_i6).plus(_dafny.ONE);
            let _155_v19;
            _155_v19 = new _dafny.CodePoint('v'.codePointAt(0));
            let _156_v20;
            _156_v20 = _module.D0.create_DC0(_155_v19);
            let _rhs30 = _156_v20;
            let _rhs31 = false;
            let _lhs21 = _133_globalState;
            _156_v20 = _rhs30;
            _lhs21.f8 = _rhs31;
            _131_v3 = _131_v3;
            let _157_v21;
            let _nw23 = Array((new BigNumber(5)).toNumber()).fill([]);
            _157_v21 = _nw23;
            let _index22 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_157_v21).length));
            (_157_v21)[_index22] = _130_v2;
            let _index23 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_157_v21).length));
            (_157_v21)[_index23] = _130_v2;
            _134_v5 = _134_v5;
          }
        }
      }
      let _158_v22;
      _158_v22 = new _dafny.CodePoint('q'.codePointAt(0));
      let _159_v23;
      _159_v23 = _module.D0.create_DC0(_158_v22);
      _159_v23 = _159_v23;
      _140_v10 = _134_v5;
      let _160_v24;
      _160_v24 = _module.D0.create_DC1(_140_v10, _128_v0, true, _128_v0);
      let _161_i7;
      _161_i7 = _dafny.ZERO;
      L5: {
        let _pat_let_tv1 = _140_v10;
        let _pat_let_tv2 = _134_v5;
        while (function (_source3) {
          if (_source3.is_DC1) {
            let _169___mcc_h0 = (_source3).cf1;
            let _170___mcc_h1 = (_source3).cf2;
            let _171___mcc_h2 = (_source3).cf3;
            let _172___mcc_h3 = (_source3).cf4;
            let _173_cf4 = _172___mcc_h3;
            let _174_cf3 = _171___mcc_h2;
            let _175_cf2 = _170___mcc_h1;
            let _176_cf1 = _169___mcc_h0;
            return _pat_let_tv1;
          } else {
            let _177___mcc_h4 = (_source3).cf0;
            let _178_cf0 = _177___mcc_h4;
            return _pat_let_tv2;
          }
        }(_160_v24)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_161_i7)) {
              break L5;
            }
            _161_i7 = (_161_i7).plus(_dafny.ONE);
            (_133_globalState).f6 = _128_v0;
            let _162_v25;
            let _nw24 = new _module.C0();
            _nw24.__ctor();
            _162_v25 = _nw24;
            let _163_v26;
            _163_v26 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_162_v25);
            (_133_globalState).f8 = ((_163_v26).Merge(_163_v26)).contains(_128_v0);
            if (_134_v5) {
              (_133_globalState).f8 = _140_v10;
              let _164_v27;
              let _nw25 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _164_v27 = _nw25;
              let _index24 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_164_v27).length));
              (_164_v27)[_index24] = _131_v3;
              let _index25 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_164_v27).length));
              let _rhs32 = _dafny.Seq.update(_module.__default.fm1(_128_v0, _140_v10, _133_globalState), _module.__default.safeIndex(_128_v0, new BigNumber((_module.__default.fm1(_128_v0, _140_v10, _133_globalState)).length)), _module.__default.fm2(_142_v12, _134_v5, _140_v10, _133_globalState));
              let _rhs33 = _140_v10;
              let _rhs34 = ((_128_v0).minus(_128_v0)).minus(new BigNumber((_131_v3).length));
              let _lhs22 = _164_v27;
              let _lhs23 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_164_v27).length));
              let _lhs24 = _133_globalState;
              let _lhs25 = _133_globalState;
              _lhs22[_lhs23] = _rhs32;
              _lhs24.f8 = _rhs33;
              _lhs25.f3 = _rhs34;
              _134_v5 = _140_v10;
              let _165_v28;
              let _nw26 = Array((new BigNumber(15)).toNumber());
              _nw26[(_dafny.ZERO).toNumber()] = _159_v23;
              _nw26[(_dafny.ONE).toNumber()] = _module.D0.create_DC0(_158_v22);
              _nw26[(new BigNumber(2)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(3)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(4)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(5)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(6)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(7)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(8)).toNumber()] = _module.D0.create_DC0(_158_v22);
              _nw26[(new BigNumber(9)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(10)).toNumber()] = _module.D0.create_DC0(_158_v22);
              _nw26[(new BigNumber(11)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(12)).toNumber()] = _module.D0.create_DC0(new _dafny.CodePoint('c'.codePointAt(0)));
              _nw26[(new BigNumber(13)).toNumber()] = _159_v23;
              _nw26[(new BigNumber(14)).toNumber()] = _159_v23;
              _165_v28 = _nw26;
              let _index26 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_165_v28).length));
              (_165_v28)[_index26] = _module.D0.create_DC0(_158_v22);
              let _index27 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_164_v27).length));
              let _index28 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_165_v28).length));
              let _rhs35 = _dafny.Seq.Concat(_131_v3, _131_v3);
              let _rhs36 = _159_v23;
              let _lhs26 = _164_v27;
              let _lhs27 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_164_v27).length));
              let _lhs28 = _165_v28;
              let _lhs29 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_165_v28).length));
              _lhs26[_lhs27] = _rhs35;
              _lhs28[_lhs29] = _rhs36;
              (_133_globalState).f7 = _128_v0;
            } else {
              _160_v24 = _160_v24;
              let _166_v29;
              let _nw27 = new _module.C0();
              _nw27.__ctor();
              _166_v29 = _nw27;
              _158_v22 = _158_v22;
              let _167_v30;
              _167_v30 = _dafny.Seq.of(_134_v5);
              (_133_globalState).f8 = !_dafny.areEqual(_167_v30, (_137_v8).fm0(_140_v10, _131_v3, _133_globalState));
              let _168_v31;
              let _nw28 = new _module.C0();
              _nw28.__ctor();
              _168_v31 = _nw28;
            }
          }
        }
      }
      let _179_v32;
      _179_v32 = _dafny.Set.fromElements(_134_v5);
      let _180_v33;
      _180_v33 = _dafny.Seq.of(_dafny.Set.fromElements(_134_v5, _140_v10, _140_v10, _140_v10, _140_v10), _179_v32, _179_v32);
      if (!_dafny.areEqual(_180_v33, _dafny.Seq.Concat(_180_v33, _dafny.Seq.of(_179_v32)))) {
        let _181_v34;
        _181_v34 = _dafny.Map.Empty.slice().updateUnsafe(_131_v3,_140_v10);
        _181_v34 = _181_v34;
        if (!((_179_v32).IsSubsetOf(_module.__default.fm3(_140_v10, _140_v10, _160_v24, _134_v5, _133_globalState)))) {
          _179_v32 = _179_v32;
          let _182_v35;
          _182_v35 = _dafny.Seq.of(_128_v0, (_dafny.ZERO).minus(_128_v0), _128_v0);
          (_133_globalState).f8 = _module.__default.fm4(_142_v12, _182_v35, false, _134_v5, _133_globalState);
          let _183_v37;
          let _init5 = ((_184_v5, _185_v0, _186_v10, _187_v12, _188_v22) => function (_189_i8) {
            return ((_module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_184_v5,_185_v0))).dtor_cf5).Merge(_dafny.Map.Empty.slice().updateUnsafe(_186_v10,new BigNumber((function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of (_dafny.Seq.of(_dafny.Seq.update(_187_v12, _module.__default.safeIndex((_dafny.ZERO).minus(_185_v0), new BigNumber((_187_v12).length)), _188_v22), _dafny.Seq.of(_188_v22))).Elements) {
                let _190_v36 = _compr_1;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.update(_187_v12, _module.__default.safeIndex((_dafny.ZERO).minus(_185_v0), new BigNumber((_187_v12).length)), _188_v22), _dafny.Seq.of(_188_v22)), _190_v36)) {
                  _coll1.push([_190_v36,_module.D0.create_DC1(!(_184_v5), _185_v0, _184_v5, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_188_v22,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length))]);
                }
              }
              return _coll1;
            }()).length)));
          })(_134_v5, _128_v0, _140_v10, _142_v12, _158_v22);
          let _nw29 = Array((new BigNumber(25)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw29.length); _i0_5++) {
            _nw29[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _183_v37 = _nw29;
          _183_v37 = _183_v37;
          let _index29 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_138_v9).length));
          (_138_v9)[_index29] = _140_v10;
          let _index30 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_138_v9).length));
          (_138_v9)[_index30] = _134_v5;
          let _index31 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_130_v2).length));
          (_130_v2)[_index31] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ngxstq"), _142_v12), _module.__default.safeIndex(_128_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ngxstq"), _142_v12)).length)), _158_v22)).length);
          let _index32 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_130_v2).length));
          (_130_v2)[_index32] = (_dafny.ZERO).minus(_128_v0);
        } else {
          let _191_v38;
          let _init6 = ((_192_v22) => function (_193_i9) {
            return _192_v22;
          })(_158_v22);
          let _nw30 = Array((new BigNumber(16)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw30.length); _i0_6++) {
            _nw30[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _191_v38 = _nw30;
          let _index33 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_191_v38).length));
          (_191_v38)[_index33] = _158_v22;
          let _index34 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_191_v38).length));
          (_191_v38)[_index34] = _158_v22;
          (_133_globalState).f6 = ((true) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.__default.fm2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), ((_194_v38) => function (_195_i10) {
            return (_194_v38)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_194_v38).length))];
          })(_191_v38)), _140_v10, (((_181_v34).contains(_142_v12)) ? ((_181_v34).get(_142_v12)) : (_module.__default.fm4(_142_v12, _dafny.Seq.of(_128_v0), _140_v10, true, _133_globalState))), _133_globalState))).length))) : (_128_v0));
          let _196_v39;
          _196_v39 = _dafny.Seq.of(_131_v3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_197_v38) => function (_198_i11) {
            return (_197_v38)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_197_v38).length))];
          })(_191_v38)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-219)), ((_199_v38) => function (_200_i12) {
            return (_199_v38)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_199_v38).length))];
          })(_191_v38))));
          _196_v39 = _dafny.Seq.of(_131_v3, _131_v3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), ((_201_v22) => function (_202_i13) {
            return _201_v22;
          })(_158_v22)), _dafny.Seq.UnicodeFromString("ngtw")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), ((_203_v38) => function (_204_i14) {
            return (_203_v38)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_203_v38).length))];
          })(_191_v38)), _142_v12);
          let _205_v40;
          _205_v40 = _dafny.Seq.of(!(_134_v5));
          (_133_globalState).f8 = _dafny.Seq.contains(_dafny.Seq.Concat(_205_v40, _205_v40), !(_179_v32).equals(_dafny.Set.fromElements(!(_140_v10))));
          let _206_v41;
          let _nw31 = Array((new BigNumber(2)).toNumber()).fill(_dafny.MultiSet.Empty);
          _206_v41 = _nw31;
          let _207_v42;
          _207_v42 = _dafny.MultiSet.fromElements(_129_v1);
          let _index35 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_206_v41).length));
          (_206_v41)[_index35] = (_207_v42).Difference(_dafny.MultiSet.fromElements(_129_v1, _129_v1));
          let _index36 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_206_v41).length));
          (_206_v41)[_index36] = _module.__default.fm5(new BigNumber(330), !(_179_v32).contains(_140_v10), (_134_v5) && (true), _133_globalState);
        }
        let _208_v43;
        let _209_v44;
        let _210_v45;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = _module.__default.m0(_133_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _208_v43 = _out6;
        _209_v44 = _out7;
        _210_v45 = _out8;
        _208_v43 = _134_v5;
        let _211_v46;
        let _nw32 = Array((new BigNumber(12)).toNumber()).fill(_module.D0.Default());
        _211_v46 = _nw32;
        let _index37 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_211_v46).length));
        (_211_v46)[_index37] = _160_v24;
        let _212_v47;
        _212_v47 = _dafny.Seq.of(_128_v0);
        let _213_v48;
        _213_v48 = _dafny.Seq.of(_module.__default.fm4(_dafny.Seq.UnicodeFromString("okbhypfhd"), _212_v47, _134_v5, _134_v5, _133_globalState));
        let _214_v49;
        _214_v49 = _dafny.Seq.of(_138_v9, _138_v9);
        let _index38 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_211_v46).length));
        let _rhs37 = ((new BigNumber((_210_v45).length)).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_213_v48)).cardinality()))).minus((((_129_v1).contains(_128_v0)) ? ((_129_v1).get(_128_v0)) : (_module.__default.fm6(new BigNumber(899), _134_v5, _133_globalState))));
        let _rhs38 = new _dafny.CodePoint('t'.codePointAt(0));
        let _rhs39 = !(_208_v43) || (_208_v43);
        let _rhs40 = _module.D0.create_DC1(_208_v43, _128_v0, _134_v5, _128_v0);
        let _rhs41 = (_214_v49)[_module.__default.safeIndex(_128_v0, new BigNumber((_214_v49).length))];
        let _lhs30 = _133_globalState;
        let _lhs31 = _211_v46;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_211_v46).length));
        _lhs30.f7 = _rhs37;
        _158_v22 = _rhs38;
        _134_v5 = _rhs39;
        _lhs31[_lhs32] = _rhs40;
        _138_v9 = _rhs41;
      } else {
        (_133_globalState).f6 = (_module.__default.fm6(_module.__default.fm6(_128_v0, _140_v10, _133_globalState), _140_v10, _133_globalState)).plus(((_134_v5) ? (_128_v0) : (new BigNumber((_dafny.Seq.UnicodeFromString("nowxxf")).length))));
        let _index39 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_138_v9).length));
        (_138_v9)[_index39] = !(_140_v10) || (_134_v5);
        let _index40 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_138_v9).length));
        (_138_v9)[_index40] = false;
        let _215_v50;
        let _init7 = ((_216_v9, _217_v5) => function (_218_i15) {
          return _dafny.Seq.of((_216_v9)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_216_v9).length))], (_216_v9)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_216_v9).length))], _217_v5);
        })(_138_v9, _134_v5);
        let _nw33 = Array((new BigNumber(22)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw33.length); _i0_7++) {
          _nw33[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _215_v50 = _nw33;
        let _index41 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_215_v50).length));
        (_215_v50)[_index41] = _dafny.Seq.of(_140_v10);
        let _219_v51;
        _219_v51 = _dafny.Seq.of(_140_v10);
        let _index42 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_215_v50).length));
        (_215_v50)[_index42] = _dafny.Seq.Concat(_219_v51, _219_v51);
        let _220_v52;
        _220_v52 = _dafny.Seq.of(_137_v8, _137_v8);
        _220_v52 = _220_v52;
        (_133_globalState).f3 = (_dafny.ZERO).minus(_128_v0);
      }
      let _221_v53;
      let _nw34 = Array((new BigNumber(8)).toNumber()).fill([]);
      _221_v53 = _nw34;
      let _222_v54;
      let _nw35 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
      _222_v54 = _nw35;
      let _223_v55;
      _223_v55 = _dafny.Seq.of(_137_v8);
      let _index43 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_222_v54).length));
      (_222_v54)[_index43] = ((_140_v10) ? (_223_v55) : (_223_v55));
      let _224_v56;
      _224_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_158_v22, _158_v22)).length),_128_v0);
      let _225_v57;
      _225_v57 = _dafny.Seq.of(_128_v0);
      let _pat_let_tv3 = _140_v10;
      let _pat_let_tv4 = _141_v11;
      let _pat_let_tv5 = _141_v11;
      let _index44 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_222_v54).length));
      let _rhs42 = function (_source4) {
        if (_source4.is_DC3) {
          let _226___mcc_h5 = (_source4).cf6;
          let _227___mcc_h6 = (_source4).cf7;
          let _228_cf7 = _227___mcc_h6;
          let _229_cf6 = _226___mcc_h5;
          return _pat_let_tv3;
        } else {
          let _230___mcc_h7 = (_source4).cf5;
          let _231_cf5 = _230___mcc_h7;
          return (_pat_let_tv4).IsDisjointFrom(_pat_let_tv5);
        }
      }(_module.__default.fm7(_134_v5, _134_v5, _133_globalState));
      let _rhs43 = _221_v53;
      let _rhs44 = _134_v5;
      let _rhs45 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_223_v55, _223_v55), _223_v55), _module.__default.safeIndex(_128_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_223_v55, _223_v55), _223_v55)).length)), _137_v8);
      let _rhs46 = ((_module.__default.fm4(_module.__default.fm1(_128_v0, _140_v10, _133_globalState), _module.__default.fm8(_140_v10, (_224_v56).update(new BigNumber(316), _128_v0), _128_v0, _133_globalState), _134_v5, _module.__default.fm4(_142_v12, _225_v57, _140_v10, false, _133_globalState), _133_globalState)) ? (_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_128_v0), _225_v57)) : (!(_134_v5)));
      let _lhs33 = _222_v54;
      let _lhs34 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_222_v54).length));
      _134_v5 = _rhs42;
      _221_v53 = _rhs43;
      _140_v10 = _rhs44;
      _lhs33[_lhs34] = _rhs45;
      _134_v5 = _rhs46;
      let _232_v58;
      let _233_v59;
      let _234_v60;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector3 = _module.__default.m0(_133_globalState);
      _out9 = _outcollector3[0];
      _out10 = _outcollector3[1];
      _out11 = _outcollector3[2];
      _232_v58 = _out9;
      _233_v59 = _out10;
      _234_v60 = _out11;
      _128_v0 = (_dafny.ZERO).minus(_128_v0);
      let _235_v61;
      let _nw36 = new _module.C0();
      _nw36.__ctor();
      _235_v61 = _nw36;
      let _236_v62;
      _236_v62 = _dafny.MultiSet.fromElements(_232_v58);
      let _237_v63;
      _237_v63 = _module.D1.create_DC3(new BigNumber((_236_v62).cardinality()), _232_v58);
      let _238_v64;
      let _nw37 = Array((new BigNumber(6)).toNumber());
      _nw37[(_dafny.ZERO).toNumber()] = _179_v32;
      _nw37[(_dafny.ONE).toNumber()] = _179_v32;
      _nw37[(new BigNumber(2)).toNumber()] = (_179_v32).Intersect((_180_v33)[_module.__default.safeIndex(_128_v0, new BigNumber((_180_v33).length))]);
      _nw37[(new BigNumber(3)).toNumber()] = _179_v32;
      _nw37[(new BigNumber(4)).toNumber()] = _179_v32;
      _nw37[(new BigNumber(5)).toNumber()] = (_179_v32).Difference(_module.__default.fm3(_140_v10, !(_134_v5), _160_v24, (_237_v63).dtor_cf7, _133_globalState));
      _238_v64 = _nw37;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_238_v64).length))) {
        let _239_i16 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_239_i16)) && ((_239_i16).isLessThan(new BigNumber((_238_v64).length))))) {
          (_238_v64)[(_239_i16)] = _179_v32;
        }
      }
      process.stdout.write(_dafny.toString(_128_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_129_v1).equals(_dafny.MultiSet.fromElements(new BigNumber(-183)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v2)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_131_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v4).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rwijw"), _dafny.Seq.UnicodeFromString("cqq")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_133_globalState).f0).equals(_dafny.MultiSet.fromElements(new BigNumber(-183)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_133_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_133_globalState).f10).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rwijw"), _dafny.Seq.UnicodeFromString("cqq")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_134_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-183),true),new BigNumber(-183)).updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-421),false),new BigNumber(183)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-421),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v9)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_140_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v11).equals(_dafny.Set.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_142_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_143_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_149_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_154_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_v22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v23).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v24).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v24).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v24).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v24).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v32).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_180_v33, _dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false), _dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_222_v54)[new BigNumber(5)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_223_v55).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v56).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),new BigNumber(-183)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_225_v57, _dafny.Seq.of(new BigNumber(-183)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_232_v58));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v59).equals(_dafny.Set.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_234_v60).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v62).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v63).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v63).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[_dafny.ZERO]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[_dafny.ONE]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[new BigNumber(2)]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[new BigNumber(3)]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[new BigNumber(4)]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v64)[new BigNumber(5)]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
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
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC3(cf6, cf7) {
      let $dt = new D1(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC2(cf5) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO, false);
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
    static create_DC5(cf9, cf10, cf11) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D2(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf8 === other.cf8;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(_dafny.ZERO, _dafny.ZERO, _module.D0.Default());
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
    static create_DC8(cf13) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = _dafny.ZERO;
      this.f6 = _dafny.ZERO;
      this.f7 = _dafny.ZERO;
      this.f8 = false;
      this._f0 = _dafny.MultiSet.Empty;
      this._f1 = [];
      this._f2 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f9 = false;
      this._f10 = _dafny.Set.Empty;
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
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
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.of(true)) : (_dafny.Seq.of(false))), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
